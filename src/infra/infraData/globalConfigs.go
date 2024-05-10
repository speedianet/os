package infraData

type globalConfigs struct {
	PkiConfDir                       string
	DomainOwnershipValidationUrlPath string
	PrimaryPublicDir                 string
	VirtualHostsConfDir              string
	MappingsConfDir                  string
}

var GlobalConfigs = globalConfigs{
	PkiConfDir:                       "/app/conf/pki",
	DomainOwnershipValidationUrlPath: "/validateOwnership",
	PrimaryPublicDir:                 "/app/html",
	VirtualHostsConfDir:              "/app/conf/nginx",
	MappingsConfDir:                  "/app/conf/nginx/mapping",
}
