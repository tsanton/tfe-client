package models

type SamlConfigurationRequestWrapper struct {
	Data SamlConfigurationRequest `json:"data"`
}

type SamlConfigurationRequest struct {
	Attributes    SamlAttributes    `json:"attributes"`
	Relationships SamlRelationships `json:"relationships"`
	Type          string            `json:"type"`
}

type SamlAttributes struct {
	AttrTeams             string `json:"attr-teams"`
	AttrUsername          string `json:"attr-username"`
	IDPCert               string `json:"idp-cert"`
	IDPMetadataURL        string `json:"idp-metadata-url"`
	IssuerURL             string `json:"issuer-url"`
	ProviderType          string `json:"provider-type"`
	SloEndpointURL        string `json:"slo-endpoint-url"`
	SsoEndpointURL        string `json:"sso-endpoint-url"`
	TeamManagementEnabled bool   `json:"team-management-enabled"`
}

type SamlRelationships struct {
	Organization SamlOrganization `json:"organization"`
}

type SamlOrganization struct {
	Data SamlOrganizationData `json:"data"`
}

type SamlOrganizationData struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}
