package tfe_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	me "github.com/tsanton/tfe-client/tfe/models/enum"
	mreq "github.com/tsanton/tfe-client/tfe/models/request"
)

func Test_service_create(t *testing.T) {
	/* Arrange */
	tfeOrganizationName := "gruntwork"
	request := mreq.Provider{
		Data: mreq.ProviderData{
			Type: "registry-providers",
			Attributes: mreq.ProviderAttributes{
				Name:         "test-provider",
				Namespace:    "gruntwork",
				RegistryName: me.RegistryTypePrivate,
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != fmt.Sprintf("/api/v2/organizations/%s/registry-providers", tfeOrganizationName) {
			t.Errorf("Request path differend from expected, got: %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("Request method differend from expected, got: %s", r.Method)
		}
		raw, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		var content mreq.Provider
		err = json.Unmarshal(raw, &content)
		if err != nil {
			panic(err)
		}

		/* Assert */
		assert.Equal(t, request.Data.Type, content.Data.Type)
		assert.Equal(t, request.Data.Attributes.Name, content.Data.Attributes.Name)
		assert.Equal(t, request.Data.Attributes.Namespace, content.Data.Attributes.Namespace)
		assert.Equal(t, request.Data.Attributes.Namespace, content.Data.Attributes.Namespace)

		resp := createProviderResponse(
			request.Data.Attributes.Name,
			request.Data.Attributes.Namespace,
			string(request.Data.Attributes.RegistryName),
			tfeOrganizationName,
			time.Now(),
		)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(resp))
	}))
	defer server.Close()

	/* Act */
	cli, _ := getTestClient(t, logger, server.URL)

	resp, err := cli.ProviderService.Create(context.Background(), tfeOrganizationName, &request)

	/* Assert */
	assert.Nil(t, err)
	assert.Equal(t, request.Data.Attributes.Name, resp.Data.Attributes.Name)
	assert.Equal(t, request.Data.Attributes.Namespace, resp.Data.Attributes.Namespace)
	assert.Equal(t, request.Data.Attributes.RegistryName, resp.Data.Attributes.RegistryName)
}

func Test_service_read(t *testing.T) {
	/* Arrange */
	tfeOrganizationName := "gruntwork"
	request := mreq.Provider{
		Data: mreq.ProviderData{
			Type: "registry-providers",
			Attributes: mreq.ProviderAttributes{
				Name:         "test-provider",
				Namespace:    "gruntwork",
				RegistryName: me.RegistryTypePrivate,
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != fmt.Sprintf("/api/v2/organizations/%s/registry-providers/%s/%s/%s", tfeOrganizationName, request.Data.Attributes.RegistryName, request.Data.Attributes.Namespace, request.Data.Attributes.Name) {
			t.Errorf("Request path differend from expected, got: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("Request method differend from expected, got: %s", r.Method)
		}
		raw, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		var content mreq.Provider
		err = json.Unmarshal(raw, &content)
		if err != nil {
			panic(err)
		}

		/* Assert */
		assert.Equal(t, request.Data.Type, content.Data.Type)
		assert.Equal(t, request.Data.Attributes.Name, content.Data.Attributes.Name)
		assert.Equal(t, request.Data.Attributes.Namespace, content.Data.Attributes.Namespace)
		assert.Equal(t, request.Data.Attributes.Namespace, content.Data.Attributes.Namespace)

		w.WriteHeader(http.StatusOK)
		resp := createProviderResponse(
			request.Data.Attributes.Name,
			request.Data.Attributes.Namespace,
			string(request.Data.Attributes.RegistryName),
			tfeOrganizationName,
			time.Now(),
		)
		_, _ = w.Write([]byte(resp))
	}))
	defer server.Close()

	/* Act */
	cli, _ := getTestClient(t, logger, server.URL)

	resp, err := cli.ProviderService.Read(context.Background(), tfeOrganizationName, &request)

	/* Assert */
	assert.Nil(t, err)
	assert.Equal(t, request.Data.Attributes.Name, resp.Data.Attributes.Name)
	assert.Equal(t, request.Data.Attributes.Namespace, resp.Data.Attributes.Namespace)
	assert.Equal(t, request.Data.Attributes.RegistryName, resp.Data.Attributes.RegistryName)
}

func Test_service_delete(t *testing.T) {
	/* Arrange */
	tfeOrganizationName := "gruntwork"
	request := mreq.Provider{
		Data: mreq.ProviderData{
			Type: "registry-providers",
			Attributes: mreq.ProviderAttributes{
				Name:         "test-provider",
				Namespace:    "gruntwork",
				RegistryName: me.RegistryTypePrivate,
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != fmt.Sprintf("/api/v2/organizations/%s/registry-providers/%s/%s/%s", tfeOrganizationName, request.Data.Attributes.RegistryName, request.Data.Attributes.Namespace, request.Data.Attributes.Name) {
			t.Errorf("Request path differend from expected, got: %s", r.URL.Path)
		}
		if r.Method != http.MethodDelete {
			t.Errorf("Request method differend from expected, got: %s", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	/* Act */
	cli, _ := getTestClient(t, logger, server.URL)

	err := cli.ProviderService.Delete(context.Background(), tfeOrganizationName, &request)

	/* Assert */
	assert.Nil(t, err)
}

// Utility function to create expected return message
func createProviderResponse(providerName, namespace, registry, organization string, timestamp time.Time) string {
	return fmt.Sprintf(`{
		"data": {
		  "id": "prov-cmEmLstBfjNNA9F3",
		  "type": "registry-providers",
		  "attributes": {
			"name": "%[1]s",
			"namespace": "%[2]s",
			"registry-name": "%[3]s",
			"created-at": "2022-02-11T19:16:59.533Z",
			"updated-at": "2022-02-11T19:16:59.533Z",
			"permissions": {
			  "can-delete": true
			}
		  },
		  "relationships": {
			"organization": {
			  "data": {
				"id": "%[4]s",
				"type": "organizations"
			  }
			},
			"versions": {
			  "data": [],
			  "links": {
				"related": "/api/v2/organizations/%[4]s/registry-providers/private/%[4]s/%[2]s"
			  }
			}
		  },
		  "links": {
			"self": "/api/v2/organizations/%[4]s/registry-providers/private/%[4]s/%[2]s"
		  }
		}
	  }`,
		providerName,
		namespace,
		registry,
		organization,
	)
}
