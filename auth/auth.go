package auth

// +k8s:deepcopy-gen=true
type InstanceToken struct {
	// Token is the jwt token identifying the loft instance.
	Token string `json:"token" query:"token" validate:"required"`
	// Certificate is the signing certificate for the token.
	Certificate string `json:"certificate" form:"certificate" query:"certificate" validate:"required"`
}
