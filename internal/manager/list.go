package manager

import "github.com/konstfish/acl-manager/internal/config"

const (
	ListFormatNetlist = "netlist"
)

func RetrieveList(conf config.ACLConfig) (string, error) {
	netList, err := downloadList(conf.List)
	if err != nil {
		return "", err
	}

	var formattedList string

	if conf.Format == ListFormatNetlist {
		formattedList = formatNetList(netList)
	}

	return formattedList, nil
}
