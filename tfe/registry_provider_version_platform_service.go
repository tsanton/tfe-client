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

type RegistryProviderVersionPlatformService struct {
	cli    *TerraformEnterpriseClient
	logger u.ILogger
}

func newRegistryProviderVersionPlatformService(cli *TerraformEnterpriseClient, logger u.ILogger) *RegistryProviderVersionPlatformService {
	return &RegistryProviderVersionPlatformService{
		cli:    cli,
		logger: logger,
	}
}

func (s *RegistryProviderVersionPlatformService) Create(ctx context.Context, organization, namespace, providerName, version string, req *mreq.ProviderVersionPlatform) (mresp.ProviderVersionPlatform, error) {
	path := fmt.Sprintf("/api/v2/organizations/%s/registry-providers/%s/%s/%s/versions/%s/platforms", organization, me.RegistryTypePrivate, namespace, providerName, version)
	resp, err := MakeRequest[*mreq.ProviderVersionPlatform, mresp.ProviderVersionPlatform](ctx, s.cli, http.MethodPost, 201, path, req)
	if err != nil {
		return mresp.ProviderVersionPlatform{}, err
	}
	return *resp, nil
}

func (s *RegistryProviderVersionPlatformService) Read(ctx context.Context, organization, namespace, providerName, version, os, arch string) (mresp.ProviderVersionPlatform, error) {
	path := fmt.Sprintf("/api/v2/organizations/%s/registry-providers/%s/%s/%s/versions/%s/platforms/%s/%s", organization, me.RegistryTypePrivate, namespace, providerName, version, os, arch)
	resp, err := MakeRequest[interface{}, mresp.ProviderVersionPlatform](ctx, s.cli, http.MethodGet, 200, path, nil)
	if err != nil {
		return mresp.ProviderVersionPlatform{}, err
	}
	return *resp, nil
}

func (s *RegistryProviderVersionPlatformService) Update(ctx context.Context) error {
	panic("not implemented")
}

func (s *RegistryProviderVersionPlatformService) Delete(ctx context.Context, organization, namespace, providerName, version, os, arch string) error {
	path := fmt.Sprintf("/api/v2/organizations/%s/registry-providers/%s/%s/%s/versions/%s/platforms/%s/%s", organization, me.RegistryTypePrivate, namespace, providerName, version, os, arch)
	_, err := MakeRequest[interface{}, interface{}](ctx, s.cli, http.MethodDelete, 204, path, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *RegistryProviderVersionPlatformService) List(ctx context.Context, organization, namespace, providerName, version string) (mresp.ProviderVersionPlatforms, error) {
	path := fmt.Sprintf("/api/v2/organizations/%s/registry-providers/%s/%s/%s/versions/%s/platforms", organization, me.RegistryTypePrivate, namespace, providerName, version)
	resp, err := MakeRequest[interface{}, mresp.ProviderVersionPlatforms](ctx, s.cli, http.MethodGet, 200, path, nil)
	if err != nil {
		return mresp.ProviderVersionPlatforms{}, err
	}
	return *resp, nil
}
