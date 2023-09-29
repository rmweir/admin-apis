package licenseapi

type ModuleName string
type FeatureName string

// Modules
const (
	KubernetesModule     ModuleName = "kubernetes"
	VirtualClusterModule ModuleName = "vcluster"
	DevPodModule         ModuleName = "devpod"
	PlatformModule       ModuleName = "platform"
)

// Features
const (
	// DevPod
	DevPod        FeatureName = "devpod"
	DevPodRunners FeatureName = "devpod-runners"

	// Virtual Clusters
	VirtualCluster                              FeatureName = "vcluster"
	VirtualClusterSleepMode                     FeatureName = "vcluster-sleep-mode"
	VirtualClusterCentralHostPathMapper         FeatureName = "vcluster-pro-distro-host-path-mapper"
	VirtualClusterProDistroImage                FeatureName = "vcluster-pro-distro-image"
	VirtualClusterProDistroAdmissionControl     FeatureName = "vcluster-pro-distro-admission-control"
	VirtualClusterProDistroBuiltInCoreDNS       FeatureName = "vcluster-pro-distro-built-in-coredns"
	VirtualClusterProDistroIsolatedControlPlane FeatureName = "vcluster-pro-distro-isolated-control-plane"
	VirtualClusterProDistroSyncPatches          FeatureName = "vcluster-pro-distro-sync-patches"

	// Spaces & Clusters
	Cluster            FeatureName = "cluster"
	ClusterAccess      FeatureName = "cluster-access"
	ClusterRoles       FeatureName = "cluster-roles"
	Namespace          FeatureName = "namespace"
	NamespaceSleepMode FeatureName = "namespace-sleep-mode"

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

	// Internal Features - not to be directly used by the license service
	SleepMode              FeatureName = "sleep-mode"
	Metrics                FeatureName = "metrics"
	ConnectLocalCluster    FeatureName = "connect-local-cluster"
	PasswordAuthentication FeatureName = "password-auth"
)
