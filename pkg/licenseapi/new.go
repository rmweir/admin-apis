package licenseapi

func New(product ProductName) *License {
	hideNamepaceFeatures := product != Loft
	hideConnectedClusterFeatures := product != VClusterPro && product != Loft
	hideVirtualClusterFeatures := product != VClusterPro && product != Loft
	hideDevPodFeatures := product != DevPodPro

	return &License{
		Modules: []*Module{
			{
				DisplayName: "Kubernetes Management",
				Name:        "kubernetes",
				Features: []*Feature{
					{
						DisplayName: "Connected Clusters",
						Name:        string(Cluster),
						Hidden:      hideConnectedClusterFeatures,
					},
					{
						DisplayName: "Cluster Access",
						Name:        string(ClusterAccess),
						Hidden:      hideConnectedClusterFeatures,
					},
					{
						DisplayName: "Cluster Role Management",
						Name:        string(ClusterRoles),
						Hidden:      hideConnectedClusterFeatures,
					},
					{
						DisplayName: "Namespace Management",
						Name:        string(Namespace),
						Hidden:      hideNamepaceFeatures,
					},
					{
						DisplayName: "Sleep Mode for Namespaces",
						Name:        string(NamespaceSleepMode),
						Hidden:      hideNamepaceFeatures,
					},
				},
			},
			{
				DisplayName: "vCluster.Pro",
				Name:        "vcluster",
				Features: []*Feature{
					{
						DisplayName: "Virtual Clusters",
						Name:        string(VirtualCluster),
						Hidden:      hideVirtualClusterFeatures,
					},
					{
						DisplayName: "Sleep Mode for Virtual Clusters",
						Name:        string(VirtualClusterSleepMode),
						Hidden:      hideVirtualClusterFeatures,
					},
					{
						DisplayName: "Security-Hardened vCluster Image",
						Name:        string(VirtualClusterProDistroImage),
						Hidden:      hideVirtualClusterFeatures,
					},
					{
						DisplayName: "Built-In CoreDNS",
						Name:        string(VirtualClusterProDistroBuiltInCoreDNS),
						Hidden:      hideVirtualClusterFeatures,
					},
					{
						DisplayName: "Virtual Admission Control",
						Name:        string(VirtualClusterProDistroAdmissionControl),
						Hidden:      true, // this feature should be restricted for now
					},
					{
						DisplayName: "Sync Patches",
						Name:        string(VirtualClusterProDistroSyncPatches),
						Hidden:      hideVirtualClusterFeatures,
					},
					{
						DisplayName: "Isolated Control Plane",
						Name:        string(VirtualClusterProDistroIsolatedControlPlane),
						Hidden:      hideVirtualClusterFeatures,
					},
					{
						DisplayName: "Central HostPath Mapper",
						Name:        string(VirtualClusterCentralHostPathMapper),
						Hidden:      hideVirtualClusterFeatures,
					},
				},
			},
			{
				DisplayName: "DevPod.Pro",
				Name:        "devpod",
				Features: []*Feature{
					{
						DisplayName: "Dev Environment Management",
						Name:        string(DevPod),
						Hidden:      hideDevPodFeatures,
					},
					{
						DisplayName: "DevPod Runners",
						Name:        string(DevPodRunners),
						Hidden:      hideDevPodFeatures,
					},
				},
			},
			{
				DisplayName: "Loft Platform",
				Name:        "platform",
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
						Hidden:      true,
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
