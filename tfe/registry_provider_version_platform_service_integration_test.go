package tfe_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	mreq "github.com/tsanton/tfe-client/tfe/models/request"
)

func Test_live_provider_service_version_platform_lifecycle(t *testing.T) {
	orgName, tfeToken := runnerValidator(t)
	cli := clientSetup(t, tfeToken)
	ctx := context.Background()

	/* Bootstrap provider service */
	version := "0.1.0"
	presp := bootstrapProvider(t, cli, orgName, "provider-version-platform-integration-test")
	gresp := boostrapGpgKey(t, cli, orgName, "provider-version-integration-test")
	pvresp := bootstrapProviderVersion(t, cli, orgName, presp.Data.Attributes.Namespace, presp.Data.Attributes.Name, version, gresp.Data.Attributes.KeyId)
	defer func() {
		perr := cli.ProviderService.Delete(ctx, orgName, string(presp.Data.Attributes.RegistryName), presp.Data.Attributes.Namespace, presp.Data.Attributes.Name)
		gerr := cli.GpgService.Delete(ctx, gresp.Data.Attributes.Namespace, gresp.Data.Attributes.KeyId)
		if perr != nil {
			panic("unable to cleanup provider")
		}
		if gerr != nil {
			panic("unable to cleanup gpg key")
		}
	}()
	request := mreq.ProviderVersionPlatform{
		Data: mreq.ProviderVersionPlatformData{
			Type: "registry-provider-version-platforms",
			Attributes: mreq.ProviderVersionPlatformDataAttributes{
				Os:      "linux",
				Arch:    "amd64",
				Shasum:  "8f69533bc8afc227b40d15116358f91505bb638ce5919712fbb38a2dec1bba38",
				Filname: "terraform-provider-provider-version-platform-integration-test_0.1.0_linux_amd64.zip",
			},
		},
	}
	/* Create provider version */
	cResp, err := cli.ProviderVersionPlatformService.Create(ctx, orgName, presp.Data.Attributes.Namespace, presp.Data.Attributes.Name, pvresp.Data.Attributes.Version, &request)
	assert.Nil(t, err)
	assert.NotNil(t, cResp)

	/* Read provider version */
	rResp, err := cli.ProviderVersionPlatformService.Read(ctx, orgName, presp.Data.Attributes.Namespace, presp.Data.Attributes.Name, pvresp.Data.Attributes.Version, cResp.Data.Attributes.Os, cResp.Data.Attributes.Arch)
	assert.Nil(t, err)
	assert.NotNil(t, rResp)

	/* List provider versions */
	lResp, err := cli.ProviderVersionPlatformService.List(ctx, orgName, presp.Data.Attributes.Namespace, presp.Data.Attributes.Name, pvresp.Data.Attributes.Version)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(lResp.Data))

	/* Delete provider version */
	err = cli.ProviderVersionPlatformService.Delete(ctx, orgName, presp.Data.Attributes.Namespace, presp.Data.Attributes.Name, pvresp.Data.Attributes.Version, cResp.Data.Attributes.Os, cResp.Data.Attributes.Arch)
	assert.Nil(t, err)

	/* List provider versions */
	lResp, err = cli.ProviderVersionPlatformService.List(ctx, orgName, presp.Data.Attributes.Namespace, presp.Data.Attributes.Name, pvresp.Data.Attributes.Version)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(lResp.Data))
}
