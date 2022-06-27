package models

type ServerError struct {
	Status int             `json:"status,omitempty"`
	Spec   ServerErrorSpec `json:"spec,omitempty"`
}

type ServerErrorSpec struct {
	Type   string `json:"type,omitempty"`
	Reason string `json:"reason,omitempty"`
}
