package tfe

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	me "github.com/tsanton/tfe-client/tfe/models/enum"
	mreq "github.com/tsanton/tfe-client/tfe/models/request"
	mresp "github.com/tsanton/tfe-client/tfe/models/response"
	u "github.com/tsanton/tfe-client/tfe/utilities"
)

type GpgService struct {
	cli    *TerraformEnterpriseClient
	logger u.ILogger
}

func newGpgService(cli *TerraformEnterpriseClient, logger u.ILogger) *GpgService {
	return &GpgService{
		cli:    cli,
		logger: logger,
	}
}

func (s *GpgService) Create(ctx context.Context, req *mreq.Gpg) (mresp.GpgKey, error) {
	path := fmt.Sprintf("/api/registry/%s/v2/gpg-keys", me.RegistryTypePrivate) //Must be private
	resp, err := MakeRequest[*mreq.Gpg, mresp.GpgKey](ctx, s.cli, http.MethodPost, 201, path, req)
	if err != nil {
		return mresp.GpgKey{}, err
	}
	return *resp, nil
}

func (s *GpgService) Read(ctx context.Context, namespace, keyId string) (mresp.GpgKey, error) {
	path := fmt.Sprintf("/api/registry/%s/v2/gpg-keys/%s/%s", string(me.RegistryTypePrivate), namespace, keyId) //Must be private
	resp, err := MakeRequest[interface{}, mresp.GpgKey](ctx, s.cli, http.MethodGet, 200, path, nil)
	if err != nil {
		return mresp.GpgKey{}, err
	}
	return *resp, nil
}

// Cannot see any usecases where change shiould not trigger replace
func (s *GpgService) Update(ctx context.Context) error {
	panic("not implemented")
}

func (s *GpgService) Delete(ctx context.Context, namespace, keyId string) error {
	path := fmt.Sprintf("/api/registry/%s/v2/gpg-keys/%s/%s", string(me.RegistryTypePrivate), namespace, keyId) //Must be private
	_, err := MakeRequest[interface{}, interface{}](ctx, s.cli, http.MethodDelete, 204, path, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *GpgService) List(ctx context.Context, namespaces []string) (mresp.GpgKeys, error) {
	path := fmt.Sprintf("/api/registry/%s/v2/gpg-keys?filter[namespace]=%s", string(me.RegistryTypePrivate), strings.Join(namespaces, ",")) //Must be private
	resp, err := MakeRequest[interface{}, mresp.GpgKeys](ctx, s.cli, http.MethodGet, 200, path, nil)
	if err != nil {
		return mresp.GpgKeys{}, err
	}
	return *resp, nil
}
