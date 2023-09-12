package licenseapi

// ResourceQuantity represents an api resource and a quantity counter for that resource type (used for limits for example).
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true
type ResourceQuantity struct {
	// Group is the api group.
	// +optional
	Group string `json:"group,omitempty"`
	// Version is the api version.
	// +optional
	Version string `json:"version,omitempty"`
	// Kind is the resource kind.
	// +optional
	Kind string `json:"kind,omitempty"`
	// Module is for display purposes.
	// +optional
	Module string `json:"module,omitempty"`
	// DisplayName is for display purposes.
	// +optional
	DisplayName string `json:"displayName,omitempty"`
	// AdjustURL contains a URL for adjusting limits.
	// +optional
	AdjustLink string `json:"adjustURL,omitempty"`
	// Quantity is the quantity for hte limit (for example).
	Quantity int64 `json:"quantity"`
}
