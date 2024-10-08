package manager

import (
	"context"
	"errors"
	"fmt"
	"time"

	netv1 "k8s.io/api/networking/v1"

	"github.com/konstfish/acl-manager/internal/config"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var ingressCache = make(map[string]config.ACLConfig)

func AddIngressToCache(conf config.ACLConfig) {
	key := fmt.Sprintf("%s/%s", conf.IngressName, conf.IngressNamespace)

	conf.Set = time.Now()

	ingressCache[key] = conf
}

func isExpired(timestamp time.Time, minutes int) bool {
	duration := time.Since(timestamp)
	println(int(duration.Minutes()))
	return duration.Minutes() > float64(minutes)
}

func GetIngressMatch(conf config.ACLConfig) (string, bool) {
	for key, cacheConf := range ingressCache {
		if conf.List == cacheConf.List &&
			conf.Type == cacheConf.Type &&
			conf.Format == cacheConf.Format &&
			conf.IngressName != cacheConf.IngressName &&
			!isExpired(cacheConf.Set, conf.Polling) {
			return key, true
		}
	}
	return "", false
}

func GetACLFromCache(ingressKey string, client client.Client) (string, error) {
	conf, ok := ingressCache[ingressKey]
	if !ok {
		return "", errors.New("Ingress not found in cache")
	}

	ingressFound := &netv1.Ingress{}

	err := client.Get(context.TODO(), types.NamespacedName{Name: conf.IngressName, Namespace: conf.IngressNamespace}, ingressFound)
	if err != nil {
		return "", err
	}

	if acl, ok := ingressFound.Annotations[conf.Destination]; ok {
		return acl, nil
	}

	return "", errors.New("ACL not found in Cache Ingress annotations")
}
