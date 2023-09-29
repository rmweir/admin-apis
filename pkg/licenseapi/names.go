package licenseapi

type ProductName string
type ModuleName string
type ResourceName string
type FeatureName string

// Products
const (
	Loft        ProductName = "loft"
	VClusterPro ProductName = "vcluster-pro"
	DevPodPro   ProductName = "devpod-pro"
)

// Modules
const (
	KubernetesModule     ModuleName = "kubernetes"
	VirtualClusterModule ModuleName = "vcluster"
	DevPodModule         ModuleName = "devpod"
	PlatformModule       ModuleName = "platform"
)

// Resources (e.g. for limits)
const (
	ConnectedCluster        ResourceName = "cluster"
	VirtualClusterInstance  ResourceName = "virtual-cluster-instance"
	SpaceInstance           ResourceName = "space-instance"
	DevPodWorkspaceInstance ResourceName = "devpod-workspace-instance"
	User                    ResourceName = "user"
)

// Features
const (
	// DevPod
	DevPod        FeatureName = "devpod"
	DevPodRunners FeatureName = "devpod-runners"

	// Virtual Clusters
	VirtualCluster                              FeatureName = "vcluster"
	VirtualClusterSleepMode                     FeatureName = "vcluster-sleep-mode"
	VirtualClusterCentralHostPathMapper         FeatureName = "vcluster-host-path-mapper"
	VirtualClusterProDistroImage                FeatureName = "vcp-distro-image"
	VirtualClusterProDistroAdmissionControl     FeatureName = "vcp-distro-admission-control"
	VirtualClusterProDistroBuiltInCoreDNS       FeatureName = "vcp-distro-built-in-coredns"
	VirtualClusterProDistroIsolatedControlPlane FeatureName = "vcp-distro-isolated-cp"
	VirtualClusterProDistroSyncPatches          FeatureName = "vcp-distro-sync-patches"

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
