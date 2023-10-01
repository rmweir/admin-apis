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

func New(product ProductName) *License {
	allowedStatus := string(FeatureStatusAllowed)

	connectedClusterStatus := string(FeatureStatusAllowed)
	if product != VClusterPro && product != Loft {
		connectedClusterStatus = string(FeatureStatusDisallowed)
	}

	namespaceStatus := string(FeatureStatusAllowed)
	if product != Loft {
		namespaceStatus = string(FeatureStatusDisallowed)
	}

	virtualClusterStatus := string(FeatureStatusAllowed)
	if product != VClusterPro && product != Loft {
		virtualClusterStatus = string(FeatureStatusDisallowed)
	}

	devpodStatus := string(FeatureStatusAllowed)
	if product != DevPodPro {
		devpodStatus = string(FeatureStatusDisallowed)
	}

	return &License{
		Modules: []*Module{
			{
				DisplayName: "Kubernetes Management",
				Name:        string(KubernetesModule),
				Limits: []*Limit{
					Limits[ConnectedCluster],
				},
				Features: []*Feature{
					{
						DisplayName: "Connected Clusters",
						Name:        string(Cluster),
						Status:      connectedClusterStatus,
					},
					{
						DisplayName: "Cluster Access",
						Name:        string(ClusterAccess),
						Status:      connectedClusterStatus,
					},
					{
						DisplayName: "Cluster Role Management",
						Name:        string(ClusterRoles),
						Status:      connectedClusterStatus,
					},
					{
						DisplayName: "Namespace Management",
						Name:        string(Namespace),
						Status:      namespaceStatus,
					},
					{
						DisplayName: "Sleep Mode for Namespaces",
						Name:        string(NamespaceSleepMode),
						Status:      namespaceStatus,
					},
				},
			},
			{
				DisplayName: "vCluster.Pro",
				Name:        string(VirtualClusterModule),
				Limits: []*Limit{
					Limits[VirtualClusterInstance],
				},
				Features: []*Feature{
					{
						DisplayName: "Virtual Clusters",
						Name:        string(VirtualCluster),
						Status:      virtualClusterStatus,
					},
					{
						DisplayName: "Sleep Mode for Virtual Clusters",
						Name:        string(VirtualClusterSleepMode),
						Status:      virtualClusterStatus,
					},
					{
						DisplayName: "Security-Hardened vCluster Image",
						Name:        string(VirtualClusterProDistroImage),
						Status:      virtualClusterStatus,
					},
					{
						DisplayName: "Built-In CoreDNS",
						Name:        string(VirtualClusterProDistroBuiltInCoreDNS),
						Status:      virtualClusterStatus,
					},
					{
						DisplayName: "Virtual Admission Control",
						Name:        string(VirtualClusterProDistroAdmissionControl),
						// this feature should be restricted for now
						Status: string(FeatureStatusDisallowed),
					},
					{
						DisplayName: "Sync Patches",
						Name:        string(VirtualClusterProDistroSyncPatches),
						Status:      virtualClusterStatus,
					},
					{
						DisplayName: "Isolated Control Plane",
						Name:        string(VirtualClusterProDistroIsolatedControlPlane),
						Status:      virtualClusterStatus,
					},
					{
						DisplayName: "Central HostPath Mapper",
						Name:        string(VirtualClusterCentralHostPathMapper),
						Status:      virtualClusterStatus,
					},
				},
			},
			{
				DisplayName: "DevPod.Pro",
				Name:        string(DevPodModule),
				Limits: []*Limit{
					Limits[DevPodWorkspaceInstance],
				},
				Features: []*Feature{
					{
						DisplayName: "Dev Environment Management",
						Name:        string(DevPod),
						Status:      devpodStatus,
					},
					{
						DisplayName: "DevPod Runners",
						Name:        string(DevPodRunners),
						Status:      devpodStatus,
					},
				},
			},
			{
				DisplayName: "Loft Platform",
				Name:        string(PlatformModule),
				Limits: []*Limit{
					Limits[User],
				},
				Features: []*Feature{
					{
						DisplayName: "Single Sign-On",
						Name:        string(SSOAuth),
						Status:      allowedStatus,
					},
					{
						DisplayName: "Multiple SSO Providers",
						Name:        string(MultipleSSOProviders),
						Status:      allowedStatus,
					},
					{
						DisplayName: "Audit Logging",
						Name:        string(AuditLogging),
						Status:      allowedStatus,
					},
					{
						DisplayName: "Automatic Auth For Ingresses",
						Name:        string(AutomaticIngressAuth),
						Status:      allowedStatus,
					},
					{
						DisplayName: "Loft as OIDC Provider",
						Name:        string(OIDCProvider),
						Status:      allowedStatus,
					},
					{
						DisplayName: "Apps",
						Name:        string(Apps),
						Status:      allowedStatus,
					},
					{
						DisplayName: "Template Versioning",
						Name:        string(TemplateVersioning),
						Status:      allowedStatus,
					},
					{
						DisplayName: "Secrets Management",
						Name:        string(Secrets),
						Status:      allowedStatus,
					},
					{
						DisplayName: "Secrets Encryption",
						Name:        string(SecretEncryption),
						Status:      allowedStatus,
					},
					{
						DisplayName: "Vault Integration",
						Name:        string(VaultIntegration),
						Status:      allowedStatus,
					},
					{
						DisplayName: "Argo Integration",
						Name:        string(ArgoIntegration),
						Status:      allowedStatus,
					},
					{
						DisplayName: "High-Availability Mode",
						Name:        string(HighAvailabilityMode),
						Status:      allowedStatus,
					},
					{
						DisplayName: "Multi-Region Mode",
						Name:        string(MultiRegionMode),
						Status:      allowedStatus,
					},
					{
						DisplayName: "Air-Gapped Mode",
						Name:        string(AirGappedMode),
						Status:      allowedStatus,
					},
					{
						DisplayName: "Custom Branding",
						Name:        string(CustomBranding),
						Status:      allowedStatus,
					},
					{
						DisplayName: "Advanced UI Customizations",
						Name:        string(AdvancedUICustomizations),
						Status:      allowedStatus,
					},
				},
			},
		},
	}
}
