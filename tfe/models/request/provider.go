package request

import (
	me "github.com/tsanton/tfe-client/tfe/models/enum"
)

type ProviderData struct {
	Type       string             `json:"type"`
	Attributes ProviderAttributes `json:"attributes"`
}

type ProviderAttributes struct {
	Name         string          `json:"name"`
	Namespace    string          `json:"namespace"`
	RegistryName me.RegistryType `json:"registry-name"`
}

/*##################
### Request body ###
##################*/

type Provider struct {
	Data ProviderData `json:"data"`
}
