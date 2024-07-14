/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"

	"github.com/konstfish/acl-manager/internal/config"
	"github.com/konstfish/acl-manager/internal/manager"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	networkingv1 "k8s.io/api/networking/v1"
)

// IngressReconciler reconciles a Ingress object
type IngressReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses/status,verbs=get;update;patch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Ingress object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.3/pkg/reconcile
func (r *IngressReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	var ingress networkingv1.Ingress

	if err := r.Get(ctx, req.NamespacedName, &ingress); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		log.Error(err, "unable to fetch Ingress")
		return ctrl.Result{}, err
	}

	log.Info("Ingress Reconcile", "Namespace", ingress.Namespace, "Name", ingress.Name)

	var conf config.ACLConfig = config.ACLConfig{
		IngressName:      ingress.Name,
		IngressNamespace: ingress.Namespace,
	}

	err := conf.ParseAnnotations(ctx, ingress.Annotations)
	if err != nil {
		log.Error(err, "Unable to parse Ingress annotations")
		return ctrl.Result{Requeue: false}, nil
	}
	if conf.List == "" {
		return ctrl.Result{Requeue: false}, nil
	}

	acl, err := manager.RetrieveList(ctx, conf, r.Client)
	if err != nil {
		log.Error(err, "Unable to retrive list")
		return ctrl.Result{Requeue: false}, nil
	}

	ingress.Annotations[conf.Destination] = acl

	// handle removal of acl-manager annotations (?)
	// delete(ingress.Annotations, destination)

	if err := r.Update(ctx, &ingress); err != nil {
		if apierrors.IsConflict(err) {
			return ctrl.Result{Requeue: true}, nil
		}
		if apierrors.IsNotFound(err) {
			return ctrl.Result{Requeue: true}, nil
		}
		log.Error(err, "unable to update Ingress")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *IngressReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&networkingv1.Ingress{}).
		Complete(r)
}
