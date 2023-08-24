package tfe

import (
	"context"
	"fmt"
	"net/http"

	u "github.com/tsanton/tfe-client/tfe/utilities"

	m "github.com/tsanton/tfe-client/tfe/models"
)

type SamlConfigurationService struct {
	cli    *TerraformEnterpriseClient
	logger u.ILogger
}

func NewSamlConfigurationService(cli *TerraformEnterpriseClient, logger u.ILogger) *SamlConfigurationService {
	return &SamlConfigurationService{
		cli:    cli,
		logger: logger,
	}
}

func (s *SamlConfigurationService) Create(ctx context.Context, organization string, config *m.SamlConfigurationRequest) (m.SamlConfigurationResponse, error) {
	path := fmt.Sprintf("/api/v2/organizations/%s/saml-configurations", organization)
	body := &m.SamlConfigurationRequestWrapper{
		Data: *config,
	}
	resp, err := MakeRequest[*m.SamlConfigurationRequestWrapper, m.SamlConfigurationResponse](ctx, s.cli, http.MethodPost, 201, path, body)
	if err != nil {
		return m.SamlConfigurationResponse{}, err
	}
	return *resp, nil
}

func (s *SamlConfigurationService) Read(ctx context.Context, id string) (m.SamlConfigurationResponseWrapper, error) {
	path := fmt.Sprintf("/api/v2/saml-configurations/%s", id)
	resp, err := MakeRequest[interface{}, m.SamlConfigurationResponseWrapper](ctx, s.cli, http.MethodGet, 200, path, nil)
	if err != nil {
		return m.SamlConfigurationResponseWrapper{}, err
	}
	return *resp, nil
}

func (s *SamlConfigurationService) Update(ctx context.Context, id string, config *m.SamlConfigurationRequest) (m.SamlConfigurationResponseWrapper, error) {
	path := fmt.Sprintf("/api/v2/saml-configurations/%s", id)
	resp, err := MakeRequest[*m.SamlConfigurationRequest, m.SamlConfigurationResponseWrapper](ctx, s.cli, http.MethodPatch, 200, path, config)
	if err != nil {
		return m.SamlConfigurationResponseWrapper{}, err
	}
	return *resp, nil
}

func (s *SamlConfigurationService) Delete(ctx context.Context, id string) error {
	path := fmt.Sprintf("/api/v2/saml-configurations/%s", id)
	_, err := MakeRequest[interface{}, interface{}](ctx, s.cli, http.MethodDelete, 204, path, nil)
	return err
}

func (s *SamlConfigurationService) List(ctx context.Context, organization string) (m.SamlConfigurationResponseListWrapper, error) {
	path := fmt.Sprintf("/api/v2/organizations/%s/saml-configurations?organization_name=%s", organization, organization)
	resp, err := MakeRequest[interface{}, m.SamlConfigurationResponseListWrapper](ctx, s.cli, http.MethodGet, 200, path, nil)
	if err != nil {
		return m.SamlConfigurationResponseListWrapper{}, err
	}
	return *resp, nil
}
