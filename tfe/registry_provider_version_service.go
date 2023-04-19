package tfe

import (
	"context"
	"fmt"
	"net/http"

	me "github.com/tsanton/tfe-client/tfe/models/enum"
	mreq "github.com/tsanton/tfe-client/tfe/models/request"
	mresp "github.com/tsanton/tfe-client/tfe/models/response"
	u "github.com/tsanton/tfe-client/tfe/utilities"
)

type RegistryProviderVersionService struct {
	cli    *TerraformEnterpriseClient
	logger u.ILogger
}

func newRegistryProviderVersionService(cli *TerraformEnterpriseClient, logger u.ILogger) *RegistryProviderVersionService {
	return &RegistryProviderVersionService{
		cli:    cli,
		logger: logger,
	}
}

func (s *RegistryProviderVersionService) Create(ctx context.Context, organization, namespace, providerName string, prov *mreq.ProviderVersion) (mresp.ProviderVersion, error) {
	path := fmt.Sprintf("/api/v2/organizations/%s/registry-providers/%s/%s/%s/versions", organization, me.RegistryTypePrivate, namespace, providerName)
	resp, err := MakeRequest[*mreq.ProviderVersion, mresp.ProviderVersion](ctx, s.cli, http.MethodPost, 201, path, prov)
	if err != nil {
		return mresp.ProviderVersion{}, err
	}
	return *resp, nil
}

func (s *RegistryProviderVersionService) Read(ctx context.Context, organization, namespace, providerName, version string) (mresp.ProviderVersion, error) {
	path := fmt.Sprintf("/api/v2/organizations/%s/registry-providers/%s/%s/%s/versions/%s", organization, me.RegistryTypePrivate, namespace, providerName, version)
	resp, err := MakeRequest[interface{}, mresp.ProviderVersion](ctx, s.cli, http.MethodGet, 200, path, nil)
	if err != nil {
		return mresp.ProviderVersion{}, err
	}
	return *resp, nil
}

func (s *RegistryProviderVersionService) Update(ctx context.Context) error {
	panic("not implemented")
}

func (s *RegistryProviderVersionService) Delete(ctx context.Context, organization, namespace, providerName, version string) error {
	path := fmt.Sprintf("/api/v2/organizations/%s/registry-providers/%s/%s/%s/versions/%s", organization, me.RegistryTypePrivate, namespace, providerName, version)
	_, err := MakeRequest[interface{}, interface{}](ctx, s.cli, http.MethodDelete, 204, path, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *RegistryProviderVersionService) List(ctx context.Context, organization, namespace, providerName string) (mresp.ProviderVersions, error) {
	path := fmt.Sprintf("/api/v2/organizations/%s/registry-providers/%s/%s/%s/versions", organization, me.RegistryTypePrivate, namespace, providerName)
	resp, err := MakeRequest[interface{}, mresp.ProviderVersions](ctx, s.cli, http.MethodGet, 200, path, nil)
	if err != nil {
		return mresp.ProviderVersions{}, err
	}
	return *resp, nil
}
