package tfe_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	api "github.com/tsanton/tfe-client/tfe"
	apim "github.com/tsanton/tfe-client/tfe/models"
	u "github.com/tsanton/tfe-client/tfe/utilities"
)

func getTestClient(t *testing.T, logger u.ILogger, host string) (*api.TerraformEnterpriseClient, error) {
	cli, err := api.NewClient(logger, &apim.ClientConfig{
		Address: host,
		Token:   "",
	})
	assert.Nil(t, err)
	return cli, err
}
