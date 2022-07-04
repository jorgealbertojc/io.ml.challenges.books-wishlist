package models

type UserAccount struct {
	ID   string           `json:"_id,omitempty" sql:"_id"`
	Spec *UserAccountSpec `json:"spec,omitempty"`
}

type UserAccountSpec struct {
	Username string `json:"username,omitempty" sql:"spe_username"`
	Password string `json:"password,omitempty" sql:"spec_password"`
}
