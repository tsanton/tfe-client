package models

type SamlConfigurationResponseWrapper struct {
	Data SamlConfigurationResponse `json:"data"`
}

type SamlConfigurationResponseListWrapper struct {
	Data []SamlConfigurationResponse `json:"data"`
}

type SamlConfigurationResponse struct {
	ID            string                     `json:"id"`
	Type          string                     `json:"type"`
	Attributes    SamlConfigurationAttribute `json:"attributes"`
	Relationships SamlConfigurationRelation  `json:"relationships"`
	Links         SamlConfigurationLink      `json:"links"`
}

type SamlConfigurationAttribute struct {
	AttrTeams             string `json:"attr-teams"`
	AttrUsername          string `json:"attr-username"`
	AuthnRequestsSigned   bool   `json:"authn-requests-signed"`
	CreatedAt             string `json:"created-at"`
	Debug                 bool   `json:"debug"`
	Enabled               bool   `json:"enabled"`
	EnabledAt             string `json:"enabled-at"`
	IdpMetadataURL        string `json:"idp-metadata-url"`
	IssuerURL             string `json:"issuer-url"`
	ProviderType          string `json:"provider-type"`
	SamlUpdatedAt         string `json:"saml-updated-at"`
	SloEndpointURL        string `json:"slo-endpoint-url"`
	SsoEndpointURL        string `json:"sso-endpoint-url"`
	SsoSessionTimeout     int    `json:"sso-session-timeout"`
	TeamManagementEnabled bool   `json:"team-management-enabled"`
	TestDetails           string `json:"test-details"`
	Tested                string `json:"tested"`
	TestedModifiedAt      string `json:"tested-modified-at"`
	WantAssertionsSigned  bool   `json:"want-assertions-signed"`
	EnabledModifiedAt     string `json:"enabled-modified-at"`
	Certificate           string `json:"certificate"`
	IdpCert               string `json:"idp-cert"`
	PrivateKey            string `json:"private-key"`
	NameID                string `json:"name-id"`
	AcsURL                string `json:"acs-url"`
	SloURL                string `json:"slo-url"`
	AudienceURL           string `json:"audience-url"`
	SpMetadataURL         string `json:"sp-metadata-url"`
}

type SamlConfigurationRelation struct {
	Organization SamlConfigurationOrganizationRelation `json:"organization"`
}

type SamlConfigurationOrganizationRelation struct {
	Data SamlConfigurationOrganizationData `json:"data"`
}

type SamlConfigurationOrganizationData struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type SamlConfigurationLink struct {
	Self       string `json:"self"`
	Enable     string `json:"enable"`
	Disable    string `json:"disable"`
	UserCounts string `json:"user-counts"`
	Test       string `json:"test"`
}
