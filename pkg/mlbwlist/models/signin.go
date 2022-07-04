package models

type Signin struct {
	ID   string      `json:"_id,omitempty"`
	Meta *SigninMeta `json:"meta"`
	Spec *SigninSpec `json:"spec,omitempty"`
}

type SigninMeta struct {
	UserID string `json:"user,omitempty"`
}

type SigninSpec struct {
	TokenHash string `json:"token,omitempty"`
}
