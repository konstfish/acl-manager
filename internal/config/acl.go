package config

import (
	"context"
	"errors"
	"strconv"

	v1 "github.com/konstfish/acl-manager/internal/apis/v1"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

type ACLConfig struct {
	IngressName      string
	IngressNamespace string
	List             string
	Type             string
	Format           string
	Destination      string
	Polling          int
}

var (
	errInvalidAnnotation = errors.New("Invalid annotation")
)

func (c *ACLConfig) ParseAnnotations(ctx context.Context, annotations map[string]string) error {
	log := log.FromContext(ctx)

	if list, ok := annotations[v1.AnnotationKeyList]; ok {
		c.List = list
	} else {
		// return early since this ingress does not have the list annotation
		return nil
	}

	if listType, ok := annotations[v1.AnnotationKeyType]; ok {
		c.Type = listType
	} else {
		if isFullURL(c.List) {
			c.Type = ListTypeHTTP
		} else if isDomain(c.List) {
			c.Type = ListTypeDNS
		} else {
			c.Type = ListTypeCM
		}
		log.Info("type not specified, auto discovered instead", "type", c.Type)
	}

	c.Format = DefaultListFormat
	if c.Type == ListTypeCM || c.Type == ListTypeSecret {
		c.Format = ListFormatCSV
	}
	if format, ok := annotations[v1.AnnotationKeyFormat]; ok {
		c.Format = format
	}

	if destination, ok := annotations[v1.AnnotationKeyDestination]; ok {
		c.Destination = destination
	} else {
		c.Destination = DefaultACLDestination
	}

	c.Polling = DefaultPollingRate
	if polling, ok := annotations[v1.AnnotationKeyPolling]; ok {
		if pollingInt, err := strconv.Atoi(polling); err == nil {
			c.Polling = pollingInt
		} else {
			log.Error(err, "invalid polling rate specified")
		}
	}

	return nil
}
