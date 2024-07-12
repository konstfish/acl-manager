package v1

const (
	// where to retrieve the ACLs content
	AnnotationKeyList = "acl-manager.konst.fish/list"

	// ACL Type (configmap, secret, list, dns, etc.)
	AnnotationKeyType = "acl-manager.konst.fish/type"

	AnnotationKeyFormat = "acl-manager.konst.fish/format"

	// ACL Destination Annotation
	AnnotationKeyDestination = "acl-manager.konst.fish/destination"
)
