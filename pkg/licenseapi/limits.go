package licenseapi

var Limits = map[ResourceName]*Limit{
	ConnectedCluster: {
		DisplayName: "Connected Clusters",
		Name:        string(ConnectedCluster),
	},
	VirtualClusterInstance: {
		DisplayName: "Virtual Clusters",
		Name:        string(VirtualClusterInstance),
	},
	DevPodWorkspaceInstance: {
		DisplayName: "Dev Environments",
		Name:        string(DevPodWorkspaceInstance),
	},
	User: {
		DisplayName: "Users",
		Name:        string(User),
	},
}
