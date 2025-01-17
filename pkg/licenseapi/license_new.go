package licenseapi

func New() *License {
	limits := make([]*Limit, 0, len(Limits))
	for _, limit := range Limits {
		limits = append(limits, limit)
	}

	return &License{
		Modules: []*Module{
			{
				DisplayName: "All Features",
				Name:        string(VirtualClusterModule),
				Limits:      limits,
				Features: []*Feature{
					{
						DisplayName: "Virtual Cluster Management",
						Name:        string(VirtualCluster),
					},
					{
						DisplayName: "Sleep Mode for Virtual Clusters",
						Name:        string(VirtualClusterSleepMode),
					},
					{
						DisplayName: "Central HostPath Mapper",
						Name:        string(VirtualClusterHostPathMapper),
					},
					{
						DisplayName: "Enterprise Plugins",
						Name:        string(VirtualClusterEnterprisePlugins),
					},
					{
						DisplayName: "Security-Hardened vCluster Image",
						Name:        string(VirtualClusterProDistroImage),
					},
					{
						DisplayName: "Built-In CoreDNS",
						Name:        string(VirtualClusterProDistroBuiltInCoreDNS),
					},
					{
						DisplayName: "Virtual Admission Control",
						Name:        string(VirtualClusterProDistroAdmissionControl),
						Status:      string(FeatureStatusHidden),
					},
					{
						DisplayName: "Sync Patches",
						Name:        string(VirtualClusterProDistroSyncPatches),
					},
					{
						DisplayName: "Embedded etcd",
						Name:        string(VirtualClusterProDistroEmbeddedEtcd),
					},
					{
						DisplayName: "Isolated Control Plane",
						Name:        string(VirtualClusterProDistroIsolatedControlPlane),
					},
					{
						DisplayName: "Centralized Admission Control",
						Name:        string(VirtualClusterProDistroCentralizedAdmissionControl),
					},
					{
						DisplayName: "Generic Sync",
						Name:        string(VirtualClusterProDistroGenericSync),
					},
					{
						DisplayName: "Translate Patches",
						Name:        string(VirtualClusterProDistroTranslatePatches),
					},
					{
						DisplayName: "KubeVirt Integration",
						Name:        string(VirtualClusterProDistroIntegrationsKubeVirt),
					},
					{
						DisplayName: "External Secrets Integration",
						Name:        string(VirtualClusterProDistroIntegrationsExternalSecrets),
					},
					{
						DisplayName: "Cert Manager Integration",
						Name:        string(VirtualClusterProDistroIntegrationsCertManager),
					},
					{
						DisplayName: "FIPS",
						Name:        string(VirtualClusterProDistroFips),
					},
					{
						DisplayName: "External Database",
						Name:        string(VirtualClusterProDistroExternalDatabase),
					},
					{
						DisplayName: "Database Connector",
						Name:        string(ConnectorExternalDatabase),
					},
					{
						DisplayName: "SleepMode",
						Name:        string(VirtualClusterProDistroSleepMode),
					},
					{
						DisplayName: "Dev Environment Management",
						Name:        string(Devpod),
					},
					{
						DisplayName: "Namespace Management",
						Name:        string(Namespaces),
					},
					{
						DisplayName: "Sleep Mode for Namespaces",
						Name:        string(NamespaceSleepMode),
					},
					{
						DisplayName: "Connected Clusters",
						Name:        string(ConnectedClusters),
					},
					{
						DisplayName: "Cluster Access",
						Name:        string(ClusterAccess),
					},
					{
						DisplayName: "Cluster Role Management",
						Name:        string(ClusterRoles),
					},
					{
						DisplayName: "Single Sign-On",
						Name:        string(SSOAuth),
					},
					{
						DisplayName: "Audit Logging",
						Name:        string(AuditLogging),
					},
					{
						DisplayName: "Automatic Auth For Ingresses",
						Name:        string(AutoIngressAuth),
					},
					{
						DisplayName: "Loft as OIDC Provider",
						Name:        string(OIDCProvider),
					},
					{
						DisplayName: "Multiple SSO Providers",
						Name:        string(MultipleSSOProviders),
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
						DisplayName: "Argo Integration",
						Name:        string(ArgoIntegration),
					},
					{
						DisplayName: "Rancher Integration",
						Name:        string(RancherIntegration),
					},
					{
						DisplayName: "Secrets Sync",
						Name:        string(Secrets),
					},
					{
						DisplayName: "Secrets Encryption",
						Name:        string(SecretEncryption),
					},
					{
						DisplayName: "HashiCorp Vault Integration",
						Name:        string(VaultIntegration),
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
