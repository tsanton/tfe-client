package response

import (
	"time"

	me "github.com/tsanton/tfe-client/tfe/models/enum"
)

type Provider struct {
	Data ProviderData `json:"data"`
}

type ProviderData struct {
	Id            string                `json:"id"`
	Type          string                `json:"type"`
	Attributes    ProviderAttributes    `json:"attributes"`
	Relationships ProviderRelationships `json:"relationships"`
	Links         ProviderLinks         `json:"links"`
}

type ProviderAttributes struct {
	Name         string              `json:"name"`
	Namespace    string              `json:"namespace"`
	RegistryName me.RegistryType     `json:"registry-name"`
	CreatedAt    time.Time           `json:"created-at"`
	UpdatedAt    time.Time           `json:"updated-at"`
	Permissions  ProviderPermissions `json:"permissions"`
}

type ProviderPermissions struct {
	CanDelete bool `json:"can-delete"`
}

type ProviderRelationships struct {
	Organization ProviderOrganizationRelationshipData `json:"organization"`
}

type ProviderOrganizationRelationshipData struct {
	Data ProviderOrganizationRelationshipDetails `json:"data"`
}

type ProviderOrganizationRelationshipDetails struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type ProviderLinks struct {
	Self string `json:"self"`
}
