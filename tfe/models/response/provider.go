package response

import (
	"time"

	me "github.com/tsanton/tfe-client/tfe/models/enum"
)

type Provider struct {
	Data providerData `json:"data"`
}

type providerData struct {
	Id            string                `json:"id"`
	Type          string                `json:"type"`
	Attributes    providerAttributes    `json:"attributes"`
	Relationships providerRelationships `json:"relationships"`
	Links         providerLinks         `json:"links"`
}

type providerAttributes struct {
	Name         string              `json:"name"`
	Namespace    string              `json:"namespace"`
	RegistryName me.RegistryType     `json:"registry-name"`
	CreatedAt    time.Time           `json:"created-at"`
	UpdatedAt    time.Time           `json:"updated-at"`
	Permissions  providerPermissions `json:"permissions"`
}

type providerPermissions struct {
	CanDelete bool `json:"can-delete"`
}

type providerRelationships struct {
	Organization providerOrganizationRelationshipData `json:"organization"`
}

type providerOrganizationRelationshipData struct {
	Data providerOrganizationRelationshipDetails `json:"data"`
}

type providerOrganizationRelationshipDetails struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type providerLinks struct {
	Self string `json:"self"`
}
