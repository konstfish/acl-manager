package manager

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func getCm(cmName string, namespace string, client client.Client) (*corev1.ConfigMap, error) {
	configMapFound := &corev1.ConfigMap{}

	err := client.Get(context.TODO(), types.NamespacedName{Name: cmName, Namespace: namespace}, configMapFound)
	if err != nil {
		return nil, err
	}

	return configMapFound, nil
}

func getCMList(name string, namespace string, client client.Client) (string, error) {
	configMapFound, err := getCm(name, namespace, client)
	if err != nil {
		return "", err
	}

	return configMapFound.Data["list"], nil
}

// secrets

func getSecret(secretName string, namespace string, client client.Client) (*corev1.Secret, error) {
	secretFound := &corev1.Secret{}

	err := client.Get(context.TODO(), types.NamespacedName{Name: secretName, Namespace: namespace}, secretFound)
	if err != nil {
		return nil, err
	}

	return secretFound, nil
}

func getSecretList(name string, namespace string, client client.Client) (string, error) {
	secretFound, err := getSecret(name, namespace, client)
	if err != nil {
		return "", err
	}

	return string(secretFound.Data["list"]), nil
}
