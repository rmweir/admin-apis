package licenseapi

func New(product ProductName) *License {
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
					},
					{
						DisplayName: "Multiple SSO Providers",
						Name:        string(MultipleSSOProviders),
					},
					{
						DisplayName: "Audit Logging",
						Name:        string(AuditLogging),
					},
					{
						DisplayName: "Automatic Auth For Ingresses",
						Name:        string(AutomaticIngressAuth),
					},
					{
						DisplayName: "Loft as OIDC Provider",
						Name:        string(OIDCProvider),
					},
					{
						DisplayName: "Apps",
						Name:        string(Apps),
					},
					{
						DisplayName: "Template Versioning",
						Name:        string(TemplateVersioning),
					},
					{
						DisplayName: "Secrets Management",
						Name:        string(Secrets),
					},
					{
						DisplayName: "Secrets Encryption",
						Name:        string(SecretEncryption),
					},
					{
						DisplayName: "Vault Integration",
						Name:        string(VaultIntegration),
					},
					{
						DisplayName: "Argo Integration",
						Name:        string(ArgoIntegration),
					},
					{
						DisplayName: "High-Availability Mode",
						Name:        string(HighAvailabilityMode),
					},
					{
						DisplayName: "Multi-Region Mode",
						Name:        string(MultiRegionMode),
					},
					{
						DisplayName: "Air-Gapped Mode",
						Name:        string(AirGappedMode),
					},
					{
						DisplayName: "Custom Branding",
						Name:        string(CustomBranding),
					},
					{
						DisplayName: "Advanced UI Customizations",
						Name:        string(AdvancedUICustomizations),
					},
				},
			},
		},
	}
}
