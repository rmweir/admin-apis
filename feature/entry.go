package feature

// Entry contains information regarding to a feature
// +k8s:deepcopy-gen=true
type Entry struct {
	Feature Feature `json:"feature"`
	// +optional
	Module string `json:"module,omitempty"`
	// +optional
	DisplayName string `json:"displayName,omitempty"`
	// +optional
	BuyLink string `json:"buy,omitempty"`
	// +optional
	TryLink string `json:"try,omitempty"`
	// +optional
	LearnMoreLink string `json:"learnMore,omitempty"`
	// +optional
	Hidden   bool `json:"hidden,omitempty"`
	Entitled bool `json:"entitled"`
	Enabled  bool `json:"enabled"`
}
