package licenseapi

// IntercomHashCreateInput is the required input data for generating a hash for a user for intercom.
// +k8s:deepcopy-gen=true
type IntercomHashCreateInput struct {
	*InstanceTokenAuth

	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

// IntercomHashCreateOutput is the struct holding all information returned from "intercom
// generate user hash" requests.
// +k8s:deepcopy-gen=true
type IntercomHashCreateOutput struct {
	CurrentTime int64  `json:"currentTime"`
	Hash        string `json:"hash"`
}
