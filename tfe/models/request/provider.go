package request

import (
	me "github.com/tsanton/tfe-client/tfe/models/enum"
)

type Provider struct {
	Data ProviderData `json:"data"`
}

type ProviderData struct {
	//Type must be 'registry-providers'
	Type       string                 `json:"type"`
	Attributes ProviderDataAttributes `json:"attributes"`
}

type ProviderDataAttributes struct {
	Name         string          `json:"name"`
	Namespace    string          `json:"namespace"`
	RegistryName me.RegistryType `json:"registry-name"`
}
