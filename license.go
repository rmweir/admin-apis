package licenseapi

// License is a struct representing the license data sent to a Loft instance after checking in with
// the license server.
// +k8s:deepcopy-gen=true
type License struct {
	// Analytics indicates the analytics endpoints and which requests should be sent to the
	// analytics server.
	// +optional
	Analytics *Analytics `json:"analytics,omitempty"`
	// DomainToken holds the JWT with the URL that the Loft instance is publicly available on.
	// (via Loft router)
	// +optional
	DomainToken string `json:"domainToken"`
	// Buttons is a slice of license server endpoints (buttons) that the Loft instance may need to
	// hit. Each Button contains the display text and link for the front end to work with.
	Buttons []Button `json:"buttons,omitempty"`
	// Announcements is a map string/string such that we can easily add any additional data without
	// needing to change types. For now, we will use the keys "name" and "content".
	// +optional
	Announcements []Announcement `json:"announcement,omitempty"`
	// BlockRequests is a slice of Request objects that the Loft instance should block from being
	// created due to license usage overrun.
	// +optional
	BlockRequests []Request `json:"blockRequests,omitempty"`
	// Limits is the number of resources allowed by the current license.
	// +optional
	Limits []ResourceQuantity `json:"limits,omitempty"`
	// Features is a map of enabled features.
	// +optional
	Features []Feature `json:"features,omitempty"`
	// IsOffline indicates if the license is an offline license or not.
	// +optional
	IsOffline bool `json:"isOffline,omitempty"`
}
