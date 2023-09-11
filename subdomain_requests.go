package licenseapi

// ClaimSubdomainInput is the required input data for "claim subdomain" operations.
// +k8s:deepcopy-gen=true
type ClaimSubdomainInput struct {
	Instance          string `json:"instance,omitempty"          form:"instance"          binding:"required"`
	Token             string `json:"token,omitempty"             form:"token"             binding:"required"`
	Subdomain         string `json:"subdomain,omitempty"         form:"subdomain"         binding:"required"`
	ReplacedSubdomain string `json:"replacedSubdomain,omitempty" form:"replacedSubdomain" binding:"required"`
	DNSTarget         string `json:"dnsTarget,omitempty"         form:"dnsTarget"         binding:"required"`
}

// ClaimSubdomainOutput is the output data for "claim subdomain" operations.
// +k8s:deepcopy-gen=true
type ClaimSubdomainOutput struct{}
