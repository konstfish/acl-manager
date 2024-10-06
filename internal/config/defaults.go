package config

// remember to update README.md when changing default values

const (
	DefaultListAuth       = ""
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
