package manager

import (
	"context"
	"encoding/base64"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func getCMList(name string, namespace string, client client.Client) (string, error) {
	configMapFound := &corev1.ConfigMap{}

	err := client.Get(context.TODO(), types.NamespacedName{Name: name, Namespace: namespace}, configMapFound)
	if err != nil {
		return "", err
	}

	return configMapFound.Data["list"], nil
}

func getSecretList(name string, namespace string, client client.Client) (string, error) {
	secretFound := &corev1.Secret{}

	err := client.Get(context.TODO(), types.NamespacedName{Name: name, Namespace: namespace}, secretFound)
	if err != nil {
		return "", err
	}

	decoded, err := base64.StdEncoding.DecodeString(string(secretFound.Data["list"]))
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}
