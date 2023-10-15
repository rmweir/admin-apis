package licenseapi

// Plan definition
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true
type Plan struct {
	// ID of the plan
	ID string `json:"id,omitempty"`

	// DisplayName is the display name of the plan
	// +optional
	DisplayName string `json:"displayName,omitempty"`

	// Status is the status of the plan (PlanStatus)
	// +optional
	Status string `json:"status,omitempty"`

	// Trial provides details about a planned, ongoing or expired trial
	// +optional
	Trial *Trial `json:"trial,omitempty"`

	// Expiration provides information about when this plan expires
	// +optional
	Expiration *PlanExpiration `json:"exp,omitempty"`

	// Features is a list of features included in the plan
	// +optional
	Features []string `json:"features,omitempty"`

	// Limits is a list of resources included in the plan and their limits
	// +optional
	Limits []Limit `json:"limits,omitempty"`

	// Quantity sets the quantity the TierResource is supposed to be at
	// +optional
	Quantity float64 `json:"quantity,omitempty"`

	// TierResource provides details about the main resource the tier quantity relates to
	// This may be nil for plans that don't have their quantity tied to a resource
	// +optional
	TierResource *PlanResource `json:"resource,omitempty"`

	// TierMode defines how tiers should be used
	// +optional
	TierMode TierMode `json:"tierMode,omitempty"`

	// Tiers is a list of tiers in this plan
	// +optional
	Tiers []PlanTier `json:"tiers,omitempty"`

	// AddOns are plans that can be added to this plan
	// +optional
	AddOns []Plan `json:"addons,omitempty"`
}

// PlanExpiration provides details about the expiration of a plan
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true
type PlanExpiration struct {
	// ExpiresAt is the unix timestamp of when the plan expires
	// +optional
	ExpiresAt int64 `json:"expiresAt,omitempty"`

	// UpgradesTo states the name of the plan that is replacing the current one upon its expiration
	// If this is nil, then this plan just expires (i.e. the subscription may be canceled, paused, etc.)
	// +optional
	UpgradesTo *string `json:"upgradesTo,omitempty"`
}

// PlanResource provides details about the main resource the tier quantity relates to
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true
type PlanResource struct {
	// Name of the resource (ResourceName)
	Name string `json:"name,omitempty"`

	// Status defines which resources will be counted towards the limit (e.g. active, total, total created etc.)
	// +optional
	Status ResourceStatus `json:"status,omitempty"`
}

// PlanTier defines a tier within a plan
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true
type PlanTier struct {
	// MinQuantity is the quantity included in this plan
	// +optional
	MinQuantity float64 `json:"min,omitempty"`

	// MaxQuantity is the max quantity that can be purchased
	// +optional
	MaxQuantity float64 `json:"max,omitempty"`

	// UnitPrice is the price per unit in this tier
	// +optional
	UnitPrice float64 `json:"unitPrice,omitempty"`

	// FlatFee is the flat fee for this tier
	// +optional
	FlatFee float64 `json:"flatFee,omitempty"`
}
