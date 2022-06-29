package models

type Signin struct {
	ID   string      `json:"_id,omitempty"`
	Spec *SigninSpec `json:"spec,omitempty"`
}

type SigninSpec struct {
	UserID    string `json:"user,omitempty"`
	TokenHash string `json:"token,omitempty"`
}
