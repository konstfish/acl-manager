package manager

const (
	ListFormatNetlist = "netlist"
)

func RetrieveList(url string, format string) (string, error) {
	netList, err := downloadList(url)
	if err != nil {
		return "", err
	}

	var formattedList string

	if format == ListFormatNetlist {
		formattedList = formatNetList(netList)
	}

	return formattedList, nil
}
