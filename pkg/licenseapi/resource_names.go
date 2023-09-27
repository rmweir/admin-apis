package licenseapi

type ResourceName string

const (
	VirtualClusters   ResourceName = "virtual-clusters"
	ConnectedClusters ResourceName = "connected-clusters"
	DevWorkspaces     ResourceName = "dev-workspaces"
	Users             ResourceName = "users"
	Teams             ResourceName = "teams"
)
