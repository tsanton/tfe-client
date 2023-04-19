package tfe_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	me "github.com/tsanton/tfe-client/tfe/models/enum"
	mreq "github.com/tsanton/tfe-client/tfe/models/request"
)

func Test_live_provider_service_lifecycle(t *testing.T) {
	orgName, token := runnerValidator(t)
	cli := clientSetup(t, token)
	ctx := context.Background()

	/* Create provider service */
	request := mreq.Provider{
		Data: mreq.ProviderData{
			Type: "registry-providers",
			Attributes: mreq.ProviderDataAttributes{
				Name:         "provider-integration-test",
				Namespace:    orgName,
				RegistryName: me.RegistryTypePrivate,
			},
		},
	}
	cResp, err := cli.ProviderService.Create(ctx, orgName, &request)
	assert.Nil(t, err)
	assert.NotNil(t, cResp)

	/* Read provider service */
	rResp, err := cli.ProviderService.Read(ctx, orgName, string(cResp.Data.Attributes.RegistryName), cResp.Data.Attributes.Namespace, cResp.Data.Attributes.Name)
	assert.Nil(t, err)
	assert.NotNil(t, rResp)

	/* Delete provider service */
	err = cli.ProviderService.Delete(ctx, orgName, string(cResp.Data.Attributes.RegistryName), cResp.Data.Attributes.Namespace, cResp.Data.Attributes.Name)
	assert.Nil(t, err)

	/* Read provider service */
	_, err = cli.ProviderService.Read(ctx, orgName, string(cResp.Data.Attributes.RegistryName), cResp.Data.Attributes.Namespace, cResp.Data.Attributes.Name)
	assert.NotNil(t, err)
}

// func Test_live_provider_service_cleanup(t *testing.T) {
// 	orgName, token := runnerValidator(t)
// 	cli := clientSetup(t, token)
// 	ctx := context.Background()
// 	err := cli.ProviderService.Delete(ctx, orgName, string(me.RegistryTypePrivate), orgName, "provider-version-platform-integration-test")
// 	assert.Nil(t, err)
// }
