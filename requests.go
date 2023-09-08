package licenseapi

// CreateInput is the required input data for "instance create" operations, that is, the
// primary endpoint that Loft instances will hit to register to the license server as well as get
// information about the instance's current license.
// +k8s:deepcopy-gen=true
type CreateInput struct {
	*InstanceTokenAuth

	// Product is the product that is being used. Can be empty, loft, devpod-pro or vcluster-pro.
	Product string `json:"product,omitempty" form:"product"`

	// Email is the admin email. Can be empty if no email is specified.
	Email string `json:"email,omitempty" form:"email"`

	LoftVersion string `json:"version"     form:"version"     validate:"required"`
	KubeVersion string `json:"kubeVersion" form:"kubeVersion" validate:"required"`

	KubeSystemNamespaceUID string `json:"kubeSystemNamespace" form:"kubeSystemNamespaceUID" validate:"required"`

	// AllocatedResources is a mapping of all resources (that we track, i.e. vcluster instances)
	// deployed on the Loft instance. The Loft instance passes this information along when it
	// performs it's checkins with the license server.
	AllocatedResources *map[string]ResourceQuantity `json:"quantities,omitempty" form:"quantities"`

	// Config is the current configuration of the Loft instance checking in.
	Config string `json:"config,omitempty" form:"config"`
}

// CreateOutput is the struct holding all information returned from "instance create"
// requests.
// +k8s:deepcopy-gen=true
type CreateOutput struct {
	// License is the license data for the requested Loft instance.
	License     *License `json:"license,omitempty"`
	CurrentTime int64    `json:"currentTime"`
}
