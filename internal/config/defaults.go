package config

const (
	DefaultACLDestination = "nginx.ingress.kubernetes.io/denylist-source-range"
	DefaultListFormat     = "netlist"
	DefaultPollingRate    = 60
)

const (
	ListFormatNetlist = "netlist"
	ListFormatCSV     = "csv"

	ListTypeHTTP   = "http"
	ListTypeDNS    = "dns"
	ListTypeCM     = "configmap"
	ListTypeSecret = "secret"
)
