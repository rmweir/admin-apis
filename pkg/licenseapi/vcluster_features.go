package licenseapi

func GetVClusterFeatures() []FeatureName {
	return []FeatureName{
		VirtualCluster,
		VirtualClusterSleepMode,
		VirtualClusterCentralHostPathMapper,
		VirtualClusterProDistroImage,
		VirtualClusterProDistroAdmissionControl,
		VirtualClusterProDistroBuiltInCoreDNS,
		VirtualClusterProDistroIsolatedControlPlane,
		VirtualClusterProDistroSyncPatches,
		VirtualClusterProDistroCentralizedAdmissionControl,
	}
}
