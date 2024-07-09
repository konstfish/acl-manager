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

	v1 "github.com/konstfish/acl-manager/internal/apis/v1"
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

	// TODO(user): your logic here

	var ingress networkingv1.Ingress
	if err := r.Get(ctx, req.NamespacedName, &ingress); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		log.Error(err, "unable to fetch Ingress")
		return ctrl.Result{}, err
	}

	log.Info("Ingress Reconcile", "Ingress", ingress)

	list, ok := ingress.Annotations[v1.AnnotationKeyList]
	if !ok {
		log.Info("Ingress does not have the required annotation", "Ingress", ingress.Name)
		return ctrl.Result{}, nil
	}

	// check for type of list (todo)
	listType, ok := ingress.Annotations[v1.AnnotationkeyType]
	if !ok {
		listType = config.DefaultListFormat
	}

	acl, err := manager.RetrieveList(list, listType)
	if err != nil {
		return ctrl.Result{}, err
	}

	// this will never be the case
	if ingress.Annotations == nil {
		ingress.Annotations = make(map[string]string)
	}

	// check the destination
	destination, ok := ingress.Annotations[v1.AnnotationKeyDestination]
	if !ok {
		destination = config.DefaultACLDestination
	}

	// append label
	ingress.Annotations[destination] = acl
	log.Info("adding label")

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

// SetupWithManager sets up the controller with the Manager.
func (r *IngressReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		// Uncomment the following line adding a pointer to an instance of the controlled resource as an argument
		For(&networkingv1.Ingress{}).
		Complete(r)
}
