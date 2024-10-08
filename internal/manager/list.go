package manager

import (
	"context"
	"errors"
	"strings"

	"github.com/konstfish/acl-manager/internal/config"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	errEmtpyList = errors.New("Empty List returned")
)

func RetrieveList(ctx context.Context, conf config.ACLConfig, client client.Client) (string, error) {
	var addressList string
	var parsedList []string
	var err error

	// retrieve list
	if conf.Type == config.ListTypeHTTP {
		// hacky, but works for now
		if conf.ListAuth != "" {
			secret, err := getSecret(conf.ListAuth, conf.IngressNamespace, client)
			if err != nil {
				return "", err
			}
			// TODO: dynamically set key
			secretString := string(secret.Data["auth"])

			conf.List = strings.Split(conf.List, "://")[0] + "://" + secretString + "@" + strings.Split(conf.List, "://")[1]
		}
		addressList, err = getHTTPList(conf.List)
	} else if conf.Type == config.ListTypeDNS {
		parsedList, err = getDNSList(conf.List)
	} else if conf.Type == config.ListTypeCM {
		addressList, err = getCMList(conf.List, conf.IngressNamespace, client)
	} else if conf.Type == config.ListTypeSecret {
		addressList, err = getSecretList(conf.List, conf.IngressNamespace, client)
	}
	if err != nil {
		return "", err
	}

	// parse list
	/// make sure there is something to parse (this is the edge case for DNS)
	if addressList != "" {
		if conf.Format == config.ListFormatNetlist {
			parsedList = praseFromNetList(addressList)
		} else if conf.Format == config.ListFormatCSV {
			parsedList = parseFromCSV(addressList)
		}
	}

	// make sure the list is not empty
	if len(parsedList) == 0 {
		return "", errEmtpyList
	}

	// format list
	formattedList := formatToCSV(parsedList)

	// validate list
	if !validateCSV(formattedList) {
		return "", errors.New("Issue formatting output")
	}

	return formattedList, nil
}
