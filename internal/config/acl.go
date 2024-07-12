package config

import (
	"errors"

	v1 "github.com/konstfish/acl-manager/internal/apis/v1"
)

type ACLConfig struct {
	List        string
	Type        string
	Format      string
	Destination string
}

var (
	errInvalidAnnotation = errors.New("Invalid annotation")
)

func (c *ACLConfig) ParseAnnotations(annotations map[string]string) error {
	if list, ok := annotations[v1.AnnotationKeyList]; ok {
		c.List = list
	} else {
		// return early since this ingress does not have the list annotation
		return nil
	}

	if listType, ok := annotations[v1.AnnotationKeyType]; ok {
		c.Type = listType
	}

	if format, ok := annotations[v1.AnnotationKeyFormat]; ok {
		c.Format = format
	} else {
		c.Format = DefaultListFormat
	}

	if destination, ok := annotations[v1.AnnotationKeyDestination]; ok {
		c.Destination = destination
	} else {
		c.Destination = DefaultACLDestination
	}

	return nil
}
