package tfe_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	m "github.com/tsanton/tfe-client/tfe/models"
)

const (
	AZURE_IDP_CERT = `MIIC6DCCAdCgAwIBAgIIFCxpKtnQBHEwDQYJKoZIhvcNAQELBQAwNDEyMDAGA1UEAwwpTWljcm9zb2Z0IEF6dXJlIEZlZGVyYXRlZCBTU08gQ2VydGlmaWNhdGUwHhcNMjMwODA3MTY1MDA2WhcNMjYwODA3MTcwMDA2WjA0MTIwMAYDVQQDDClNaWNyb3NvZnQgQXp1cmUgRmVkZXJhdGVkIFNTTyBDZXJ0aWZpY2F0ZTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAMH0tecNjhyGghiSbnCQ6mWUw83TLTuXm5B4VeVsN44tLw9UcH0X1Qu1Zp1kFXkfV4Sb+zusfywADpRku/V3DoSzVi0tIMt/X39NqTihbscX3M5Ridg1supHU8qwwv0AUC+raXe96mQXP5UmB6F+ChrWMEEXmGLN8cTvxTdnAE7Nos6/9rfhoOiS1ZhPDBrg8gyKfMZXzDOV/OooEUs5U7pDYXPpFSc+VdZAopwan2JxjcvuV/XYMzr8MRSL7Bl1TwivoHe/hs2xEdMwYfW9/LflZNnOrnAtUybZrlWRjpgZYe82u0enlAOVmqZxJ35AoinU2b1UYp75A2V3VskqUWsCAwEAATANBgkqhkiG9w0BAQsFAAOCAQEANnUkGZf1DWyFIhOQos/6KfEJqsAvkaSL5jC0w9jccFWU/QPoAOI4BKWBEhyYgqt7kVLcW2NThV8aJ5nHjH2l4nQGARt8hCvn8xIZQ001ezYcMyEzkEMGArnOFr0Xv/1gPRYJJsPzXP0rdNcU3RA8OVDy7VgGAWyYW1PoKJuYfvXPkBSkjyUULPfBMtf35663gZWrEzKiEA+IdTwC+B4IZnSg8UHMGwihZdK1c/bh7awD/f1vPYQ6FMxbno/ZZngeAd+F/9OXLjK/qHLDx8kyLV4mi1Go7K/H94VS+UQNfTCPdX4Ngo1t9A+kpLkUY0Kyg+V3ZLUirKDG0XP/Vu9W5g==`
)

func Test_live_saml_config_lifecycle(t *testing.T) {

	orgName, token := runnerValidator(t)
	cli := clientSetup(t, token)
	ctx := context.Background()

	request := m.SamlConfigurationRequest{
		Attributes: m.SamlAttributes{
			AttrTeams:             "MemberOf",
			AttrUsername:          "Username",
			IDPCert:               AZURE_IDP_CERT,
			IDPMetadataURL:        "",
			IssuerURL:             "foobar",
			ProviderType:          "azure",
			SloEndpointURL:        "",
			SsoEndpointURL:        "test",
			TeamManagementEnabled: true,
		},
		Relationships: m.SamlRelationships{
			Organization: m.SamlOrganization{
				Data: m.SamlOrganizationData{
					Type: "organizations",
					ID:   orgName,
				},
			},
		},
		Type: "saml-configurations",
	}
	cResp, err := cli.SamlConfigurationService.Create(ctx, orgName, &request)
	assert.Nil(t, err)
	assert.NotNil(t, cResp)

	// /* Read provider service */
	_, err = cli.SamlConfigurationService.List(ctx, orgName)
	assert.Nil(t, err)
	rResp, err := cli.SamlConfigurationService.Read(ctx, orgName)
	assert.Nil(t, err)
	assert.NotNil(t, rResp)

	/* Delete provider service */
	err = cli.SamlConfigurationService.Delete(ctx, cResp[0].ID)
	assert.Nil(t, err)

	/* Read provider service */
	_, err = cli.SamlConfigurationService.List(ctx, orgName)
	assert.Nil(t, err)
	_, err = cli.SamlConfigurationService.Read(ctx, "samlconf-9CDto22uBFFoocweGLDqjijuj359mm")
	assert.Nil(t, err)
}
