package licenseapi

// ButtonActionInput defines the payload that needs to be sent to a button's action URL
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true
type ButtonActionInput struct {
	*InstanceTokenAuth `json:",omitempty" hash:"-"`

	// URL is the url that the button leads to
	URL string `json:"url"`

	// +optional
	ReturnURL string `json:"returnURL,omitempty"`
}

// ButtonActionOutput specifies the response
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true
type ButtonActionOutput struct {
	// RedirectURL to redirect the user to.
	// +optional
	RedirectURL string `json:"redirectURL,omitempty"`

	// HTML to display to the user.
	// +optional
	HTML string `json:"html,omitempty"`

	// Buttons to be shown to the user alongside other content (e.g. HTML).
	// +optional
	Buttons []Button `json:"buttons,omitempty"`
}
