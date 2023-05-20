package registry

type Registration struct {
	ServiceName      ServiceName
	ServiceURL       string
	RequiredServices []ServiceName
	ServiceUpdateURL string
	HeartbeatURL     string
}

type ServiceName string

const (
	RegistrySpokeService = ServiceName("RegistryService")
	LogService           = ServiceName("LogService")
	EmailService         = ServiceName("EmailService")
	TideService          = ServiceName("TideService")
	MarinaService        = ServiceName("MarinaService")
)

type patchEntry struct {
	Name ServiceName
	URL  string
}

type patch struct {
	Added   []patchEntry
	Removed []patchEntry
}
