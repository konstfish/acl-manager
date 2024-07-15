package v1

const (
	// where to retrieve the ACLs content
	AnnotationKeyList = "acl-manager.konst.fish/list"

	// Type (configmap, secret, list, dns)
	AnnotationKeyType = "acl-manager.konst.fish/type"

	// Format (netlist, csv)
	AnnotationKeyFormat = "acl-manager.konst.fish/format"

	// Destination Annotation
	AnnotationKeyDestination = "acl-manager.konst.fish/destination"

	// Polling interval (in minutes)
	AnnotationKeyPolling = "acl-manager.konst.fish/polling"
)
