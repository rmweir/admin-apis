package licenseapi

var Limits = map[LimitName]*Limit{
	ConnectedClusterLimit: {
		DisplayName: "Connected Clusters",
		Name:        string(ConnectedClusterLimit),
	},
	VirtualClusterLimit: {
		DisplayName: "Virtual Clusters",
		Name:        string(VirtualClusterLimit),
	},
	DevEnvironmentLimit: {
		DisplayName: "Dev Environments",
		Name:        string(DevEnvironmentLimit),
	},
	UserLimit: {
		DisplayName: "Users",
		Name:        string(UserLimit),
	},
}
