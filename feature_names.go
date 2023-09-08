package licenseapi

type FeatureName string

var (
	AllFeatures FeatureName = "all"

	// DevPod
	DevPodWorkspaces FeatureName = "devpod-workspaces"

	// Runners
	Runners FeatureName = "runners"

	// Virtual Clusters
	VirtualClusters FeatureName = "vcluster"

	VirtualClusterAdmissionControl      FeatureName = "vcluster-admission-control"
	VirtualClusterBuiltInCoreDNS        FeatureName = "vcluster-built-in-coredns"
	VirtualClusterCentralHostPathMapper FeatureName = "vcluster-host-path-mapper"
	VirtualClusterIsolatedControlPlane  FeatureName = "vcluster-isolated-control-plane"
	VirtualClusterSleepMode             FeatureName = "vcluster-sleep-mode"
	VirtualClusterSyncPatches           FeatureName = "vcluster-sync-patches"

	// Spaces & Clusters
	ClusterAccess     FeatureName = "cluster-access"
	ClusterRoles      FeatureName = "cluster-roles"
	ConnectedClusters FeatureName = "clusters"
	Spaces            FeatureName = "spaces"
	SpacesSleepMode   FeatureName = "spaces-sleep-mode"

	// Auth-Related Features
	AuditLogging         FeatureName = "audit-logging"
	AutomaticIngressAuth FeatureName = "auto-ingress-authentication"
	MultipleSSOProviders FeatureName = "multiple-sso-providers"
	OIDCProvider         FeatureName = "oidc-provider"
	SSOAuth              FeatureName = "sso-authentication"

	// Templating Features
	Apps               FeatureName = "apps"
	TemplateVersioning FeatureName = "template-versioning"

	// Secrets
	Secrets          FeatureName = "secrets"
	SecretEncryption FeatureName = "secret-encryption"

	// Integrations
	ArgoIntegration  FeatureName = "argo-integration"
	VaultIntegration FeatureName = "vault-integration"

	// HA & Other Advanced Deployment Features
	AirGappedMode        FeatureName = "air-gapped-mode"
	HighAvailabilityMode FeatureName = "ha-mode"
	MultiRegionMode      FeatureName = "multi-region-mode"

	// UI Customization Features
	AdvancedUICustomizations FeatureName = "advanced-ui-customizations"
	CustomBranding           FeatureName = "custom-branding"
)
