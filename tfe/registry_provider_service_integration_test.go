package tfe_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	api "github.com/tsanton/tfe-client/tfe"
	apim "github.com/tsanton/tfe-client/tfe/models"
	me "github.com/tsanton/tfe-client/tfe/models/enum"
	mreq "github.com/tsanton/tfe-client/tfe/models/request"
	u "github.com/tsanton/tfe-client/tfe/utilities"
)

func Test_live_provider_service_lifecycle(t *testing.T) {
	tfeOrgName := u.GetEnv("TFE_ORG_NAME", "")
	tfeToken := u.GetEnv("TFE_TOKEN", "")
	run := u.GetEnv("TFE_RUN_LIVE_TESTS", false)
	if !run && tfeOrgName != "" && tfeToken != "" {
		t.Skip("Skipping test 'Test_live_provider_service_lifecycle'")
	}
	cli, err := api.NewClient(logger, &apim.ClientConfig{
		Address: "https://app.terraform.io",
		Token:   tfeToken,
	})
	if err != nil {
		t.Errorf("Error creating client: %s", err)
		t.FailNow()
	}

	ctx := context.Background()

	/* List provicer services */

	/* Create provider service */
	request := mreq.Provider{
		Data: mreq.ProviderData{
			Type: "registry-providers",
			Attributes: mreq.ProviderAttributes{
				Name:         "test-provider",
				Namespace:    tfeOrgName,
				RegistryName: me.RegistryTypePrivate,
			},
		},
	}
	cResp, err := cli.ProviderService.Create(ctx, tfeOrgName, &request)
	assert.Nil(t, err)
	assert.NotNil(t, cResp)

	/* Read provider service */
	rResp, err := cli.ProviderService.Read(ctx, tfeOrgName, &request)
	assert.Nil(t, err)
	assert.NotNil(t, rResp)

	/* Delete provider service */
	err = cli.ProviderService.Delete(ctx, tfeOrgName, &request)
	assert.Nil(t, err)

	/* Read provider service */
	_, err = cli.ProviderService.Read(ctx, tfeOrgName, &request)
	assert.NotNil(t, err)
}

// func Test_live_provider_service_cleanup(t *testing.T) {
// 	tfeOrgName := u.GetEnv("TFE_ORG_NAME", "")
// 	tfeToken := u.GetEnv("TFE_TOKEN", "")

// 	cli, err := api.NewClient(logger, &apim.ClientConfig{
// 		Address: "https://app.terraform.io",
// 		Token:   tfeToken,
// 	})
// 	if err != nil {
// 		t.Errorf("Error creating client: %s", err)
// 		t.FailNow()
// 	}

// 	ctx := context.Background()
// 	request := mreq.Provider{
// 		Data: mreq.ProviderData{
// 			Type: "registry-providers",
// 			Attributes: mreq.ProviderAttributes{
// 				Name:         "demo-provider",
// 				Namespace:    tfeOrgName,
// 				RegistryName: me.RegistryTypePrivate,
// 			},
// 		},
// 	}
// 	err = cli.ProviderService.Delete(ctx, tfeOrgName, &request)
// 	assert.Nil(t, err)
// }
