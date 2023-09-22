package license

type FeatureName string

const (
	// DevPod
	DevPod        FeatureName = "devpod"
	DevPodRunners FeatureName = "devpod-runners"

	// Virtual Clusters
	VirtualCluster                      FeatureName = "vcluster"
	VirtualClusterSleepMode             FeatureName = "vcluster-sleep-mode"
	VirtualClusterAdmissionControl      FeatureName = "vcluster-admission-control"
	VirtualClusterBuiltInCoreDNS        FeatureName = "vcluster-built-in-coredns"
	VirtualClusterCentralHostPathMapper FeatureName = "vcluster-host-path-mapper"
	VirtualClusterIsolatedControlPlane  FeatureName = "vcluster-isolated-control-plane"
	VirtualClusterSyncPatches           FeatureName = "vcluster-sync-patches"

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

	// Utils - not to be directly used by the license service
	Metrics                FeatureName = "metrics"
	ConnectLocalCluster    FeatureName = "connect-local-cluster"
	PasswordAuthentication FeatureName = "password-auth"
)
