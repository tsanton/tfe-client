package tfe_test

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	api "github.com/tsanton/tfe-client/tfe"
	apim "github.com/tsanton/tfe-client/tfe/models"
	me "github.com/tsanton/tfe-client/tfe/models/enum"
	mreq "github.com/tsanton/tfe-client/tfe/models/request"
	mresp "github.com/tsanton/tfe-client/tfe/models/response"
	u "github.com/tsanton/tfe-client/tfe/utilities"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
)

func getTestClient(t *testing.T, logger u.ILogger, host string) (*api.TerraformEnterpriseClient, error) {
	cli, err := api.NewClient(logger, &apim.ClientConfig{
		Address: host,
		Token:   "",
	})
	assert.Nil(t, err)
	return cli, err
}

func runnerValidator(t *testing.T) (string, string) {
	orgName := u.GetEnv("TFE_ORG_NAME", "")
	token := u.GetEnv("TFE_TOKEN", "")
	run := u.GetEnv("TFE_RUN_LIVE_TESTS", false)
	if !run && orgName != "" && token != "" {
		t.Skip("Skipping test 'Test_live_provider_service_version_lifecycle'")
	}
	return orgName, token
}

func clientSetup(t *testing.T, token string) *api.TerraformEnterpriseClient {
	cli, err := api.NewClient(logger, &apim.ClientConfig{
		Address: "https://app.terraform.io",
		Token:   token,
	})
	if err != nil {
		t.Errorf("Error creating client: %s", err)
		t.FailNow()
	}
	return cli
}

func boostrapGpgKey(t *testing.T, cli *api.TerraformEnterpriseClient, orgName, comment string) *mresp.GpgKey {
	entity, err := openpgp.NewEntity(orgName, comment, "donotreply@gruntwork.com", &packet.Config{RSABits: 4096})
	if err != nil {
		panic("unable to generate GPG entity in boostrapGpgKey")
	}

	/* Generate GPG key */
	publicKeyString, err := generateGpgKey(entity)
	if err != nil {
		panic("unable to generate GPG key in boostrapGpgKey")
	}

	/* Create GPG key*/
	request := &mreq.Gpg{
		Data: mreq.GpgData{
			Type: "gpg-keys",
			Attributes: mreq.GpgDataAttributes{
				AsciiArmor: publicKeyString,
				Namespace:  orgName,
			},
		},
	}
	cResp, err := cli.GpgService.Create(context.Background(), request)
	if err != nil {
		panic("unable to boostrap gpg key")
	}
	return &cResp
}

func bootstrapProvider(t *testing.T, cli *api.TerraformEnterpriseClient, orgName, providerName string) *mresp.Provider {
	request := mreq.Provider{
		Data: mreq.ProviderData{
			Type: "registry-providers",
			Attributes: mreq.ProviderDataAttributes{
				Name:         providerName,
				Namespace:    orgName,
				RegistryName: me.RegistryTypePrivate,
			},
		},
	}
	cResp, err := cli.ProviderService.Create(context.Background(), orgName, &request)
	if err != nil {
		panic("unable to boostrap provider")
	}
	return &cResp
}

func bootstrapProviderVersion(t *testing.T, cli *api.TerraformEnterpriseClient, orgName, namespace, providerName, version, keyId string) *mresp.ProviderVersion {
	request := mreq.ProviderVersion{
		Data: mreq.ProviderVersionData{
			Type: "registry-provider-versions",
			Attributes: mreq.ProviderVersionDataAttributes{
				Version:   "0.1.0",
				KeyId:     keyId,
				Protocols: []string{"5.0", "6.0"},
			},
		},
	}
	/* Create provider version */
	cResp, err := cli.ProviderVersionService.Create(context.Background(), orgName, namespace, providerName, &request)
	if err != nil {
		panic("unable to boostrap provider version")
	}
	return &cResp
}

func generateGpgKey(entity *openpgp.Entity) (string, error) {
	var publicKeyBuf bytes.Buffer
	err := entity.Serialize(&publicKeyBuf)
	if err != nil {
		fmt.Println("Error serializing public key:", err)
		return "", err
	}

	// Convert the public key to an armored string
	publicKeyArmorBuf := bytes.Buffer{}
	w, err := armor.Encode(&publicKeyArmorBuf, "PGP PUBLIC KEY BLOCK", nil)
	if err != nil {
		fmt.Println("Error encoding public key:", err)
		return "", err
	}
	_, err = w.Write(publicKeyBuf.Bytes())
	if err != nil {
		fmt.Println("Error writing public key to armored buffer:", err)
		return "", err
	}
	w.Close()

	return publicKeyArmorBuf.String(), nil
}
