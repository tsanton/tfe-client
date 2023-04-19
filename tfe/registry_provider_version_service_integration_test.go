package tfe_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	mreq "github.com/tsanton/tfe-client/tfe/models/request"
)

func Test_live_provider_service_version_lifecycle(t *testing.T) {
	orgName, tfeToken := runnerValidator(t)
	cli := clientSetup(t, tfeToken)
	ctx := context.Background()

	/* Bootstrap provider service */
	presp := bootstrapProvider(t, cli, orgName, "provider-version-integration-test")
	gresp := boostrapGpgKey(t, cli, orgName, "provider-version-integration-test")
	defer func() {
		perr := cli.ProviderService.Delete(ctx, orgName, string(presp.Data.Attributes.RegistryName), presp.Data.Attributes.Namespace, presp.Data.Attributes.Name)
		gerr := cli.GpgService.Delete(ctx, gresp.Data.Attributes.Namespace, gresp.Data.Attributes.KeyId)
		if perr != nil {
			panic("unable to cleanup provider service")
		}
		if gerr != nil {
			panic("unable to cleanup gpg key")
		}
	}()
	request := mreq.ProviderVersion{
		Data: mreq.ProviderVersionData{
			Type: "registry-provider-versions",
			Attributes: mreq.ProviderVersionDataAttributes{
				Version:   "0.1.0",
				KeyId:     gresp.Data.Attributes.KeyId,
				Protocols: []string{"5.0", "6.0"},
			},
		},
	}
	/* Create provider version */
	cResp, err := cli.ProviderVersionService.Create(ctx, orgName, presp.Data.Attributes.Namespace, presp.Data.Attributes.Name, &request)
	assert.Nil(t, err)
	assert.NotNil(t, cResp)

	/* Read provider version */
	rResp, err := cli.ProviderVersionService.Read(ctx, orgName, presp.Data.Attributes.Namespace, presp.Data.Attributes.Name, cResp.Data.Attributes.Version)
	assert.Nil(t, err)
	assert.NotNil(t, rResp)

	/* List provider versions */
	lResp, err := cli.ProviderVersionService.List(ctx, orgName, presp.Data.Attributes.Namespace, presp.Data.Attributes.Name)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(lResp.Data))

	/* Delete provider version */
	err = cli.ProviderVersionService.Delete(ctx, orgName, presp.Data.Attributes.Namespace, presp.Data.Attributes.Name, cResp.Data.Attributes.Version)
	assert.Nil(t, err)

	/* List provider versions */
	lResp, err = cli.ProviderVersionService.List(ctx, orgName, presp.Data.Attributes.Namespace, presp.Data.Attributes.Name)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(lResp.Data))
}
