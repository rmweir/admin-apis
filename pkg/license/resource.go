package license

// ResourceQuantity represents an api resource and a quantity counter for that resource type (used for limits for example).
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true
type ResourceQuantity struct {
	// Group is the api group.
	// +optional
	Name string `json:"name,omitempty"`
	// DisplayName is for display purposes.
	// +optional
	DisplayName string `json:"displayName,omitempty"`
	// AdjustButton is the button to be shown for adjusting the limit (e.g. buying more seats)
	// +optional
	AdjustButton Button `json:"adjustButton,omitempty"`
	// Limit specifies the limit for this resource.
	// +optional
	Limit int64 `json:"limit"`
	// Usage specifies the amount of resources currently existing.
	// +optional
	Usage int64 `json:"usage"`

	// BlockRequests specifies which requests the product should block when a limit is exceeded.
	// +optional
	BlockRequests []Request `json:"blockRequests,omitempty"`
}
