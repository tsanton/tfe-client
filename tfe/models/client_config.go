package models

type ClientConfig struct {
	// The address of the Terraform Enterprise API.
	Address string

	// API token used to access the Terraform Enterprise API.
	Token string
}
