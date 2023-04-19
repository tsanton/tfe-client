package tfe_test

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"testing"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"

	"github.com/stretchr/testify/assert"
	areq "github.com/tsanton/tfe-client/tfe/models/request"
)

func Test_live_gpg_key_lifecycle(t *testing.T) {
	orgName, token := runnerValidator(t)
	cli := clientSetup(t, token)
	ctx := context.Background()

	entity, err := openpgp.NewEntity("Gruntwork", "Integration test GPG key", "donotreply@gruntwork.com", &packet.Config{RSABits: 4096})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating entity: %v\n", err)
		os.Exit(1)
	}

	/* Generate GPG key */
	publicKeyString, err := generateGpgKey(entity)
	if err != nil {
		t.FailNow()
	}

	/* List GPG keys */
	origKeys, err := cli.GpgService.List(ctx, []string{orgName})
	assert.Nil(t, err)
	t.Logf("Number keys initial keys: %d", len(origKeys.Data))
	// assert.Equal(t, 0, len(keys.Data))

	/* Create GPG key*/
	cReq := &areq.Gpg{
		Data: areq.GpgData{
			Type: "gpg-keys",
			Attributes: areq.GpgDataAttributes{
				AsciiArmor: publicKeyString,
				Namespace:  orgName,
			},
		},
	}
	cResp, err := cli.GpgService.Create(ctx, cReq)
	assert.Nil(t, err)

	/* Read GPG key */
	rResp, err := cli.GpgService.Read(ctx, cReq.Data.Attributes.Namespace, cResp.Data.Attributes.KeyId)
	assert.Nil(t, err)
	assert.NotNil(t, rResp)

	/* List GPG Keys */
	keys, err := cli.GpgService.List(ctx, []string{orgName})
	assert.Nil(t, err)
	t.Logf("Number keys after create: %d", len(keys.Data))

	/* Delete GPG key */
	err = cli.GpgService.Delete(ctx, cReq.Data.Attributes.Namespace, cResp.Data.Attributes.KeyId)
	assert.Nil(t, err)

	/* Assert key deleted */
	finalKeys, err := cli.GpgService.List(ctx, []string{orgName})
	assert.Nil(t, err)
	t.Logf("Number keys after delete: %d", len(finalKeys.Data))
	assert.Equal(t, len(origKeys.Data), len(finalKeys.Data))

	/* Cleanup all keys */
	// for _, key := range finalKeys.Data {
	// 	err = cli.GpgService.Delete(ctx, key.Attributes.Namespace, key.Attributes.KeyId)
	// 	if err != nil {
	// 		panic("whops")
	// 	}
	// }
}

func Test_gpg_key(t *testing.T) {
	entity, err := openpgp.NewEntity("Gruntwork", "Integration test GPG key", "donotreply@gruntwork.com", &packet.Config{RSABits: 4096})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating entity: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Generated public key:")
	publicKeyString, err := generateGpgKey(entity)
	if err != nil {
		t.FailNow()
	}

	publicKeyReader := bytes.NewBufferString(publicKeyString)
	_, err = openpgp.ReadArmoredKeyRing(publicKeyReader)
	assert.Nil(t, err)
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
