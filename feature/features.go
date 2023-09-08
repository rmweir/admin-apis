package feature

type Feature string

var (
	AllFeatures Feature = "all"

	// DevPod
	DevPodWorkspaces Feature = "devpod-workspaces"

	// Runners
	Runners Feature = "runners"

	// Virtual Clusters
	VirtualClusters Feature = "vcluster"

	VirtualClusterAdmissionControl      Feature = "vcluster-admission-control"
	VirtualClusterBuiltInCoreDNS        Feature = "vcluster-built-in-coredns"
	VirtualClusterCentralHostPathMapper Feature = "vcluster-host-path-mapper"
	VirtualClusterIsolatedControlPlane  Feature = "vcluster-isolated-control-plane"
	VirtualClusterSleepMode             Feature = "vcluster-sleep-mode"
	VirtualClusterSyncPatches           Feature = "vcluster-sync-patches"

	// Spaces & Clusters
	ClusterAccess     Feature = "cluster-access"
	ClusterRoles      Feature = "cluster-roles"
	ConnectedClusters Feature = "clusters"
	Spaces            Feature = "spaces"
	SpacesSleepMode   Feature = "spaces-sleep-mode"

	// Auth-Related Features
	AuditLogging         Feature = "audit-logging"
	AutomaticIngressAuth Feature = "auto-ingress-authentication"
	MultipleSSOProviders Feature = "multiple-sso-providers"
	OIDCProvider         Feature = "oidc-provider"
	SSOAuth              Feature = "sso-authentication"

	// Templating Features
	Apps               Feature = "apps"
	TemplateVersioning Feature = "template-versioning"

	// Secrets
	Secrets          Feature = "secrets"
	SecretEncryption Feature = "secret-encryption"

	// Integrations
	ArgoIntegration  Feature = "argo-integration"
	VaultIntegration Feature = "vault-integration"

	// HA & Other Advanced Deployment Features
	AirGappedMode        Feature = "air-gapped-mode"
	HighAvailabilityMode Feature = "ha-mode"
	MultiRegionMode      Feature = "multi-region-mode"

	// UI Customization Features
	AdvancedUICustomizations Feature = "advanced-ui-customizations"
	CustomBranding           Feature = "custom-branding"
)
