package licenseapi

// ResourceQuantity represents an api resource and a quantity counter for that resource type (used for limits for example).
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true
type ResourceLimit struct {
	// Name is the name of the resource.
	// +optional
	Name string `json:"name,omitempty"`
	// DisplayName is for display purposes.
	// +optional
	DisplayName string `json:"displayName,omitempty"`
	// AdjustButton is the button to be shown for adjusting the limit (e.g. buying more seats)
	// +optional
	AdjustButton *Button `json:"adjustButton,omitempty"`

	// Limit specifies the limit for this resource.
	// +optional
	Limit *ResourceCount `json:",omitempty"`
}

// ResourceCount stores the number of existing, active and total number of resources created.
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true
type ResourceCount struct {
	// Active specifies the number of currently active resource (non-sleeping).
	// +optional
	Active int64 `json:"active,omitempty"`
	// Total specifies the number of currently existing resources.
	// +optional
	Total int64 `json:"total,omitempty"`
	// TotalCreated is a continuous counter of the amount of resources ever created.
	// +optional
	TotalCreated int64 `json:"created,omitempty"`
}
