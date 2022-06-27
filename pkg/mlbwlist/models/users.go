package models

type UserAccount struct {
	ID        string           `json:"_id,omitempty" sql:"_id"`
	CreatedAt int              `json:"created_at,omitempty" sql:"created_at"`
	Spec      *UserAccountSpec `json:"spec,omitempty"`
}

type UserAccountSpec struct {
	Username string `json:"username,omitempty" sql:"spe_username"`
	Password string `json:"password,omitempty" sql:"spec_password"`
}
