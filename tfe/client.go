package tfe

import (
	"net/http"
	"net/url"

	"github.com/hashicorp/go-cleanhttp"
	u "github.com/tsanton/tfe-client/tfe/utilities"

	m "github.com/tsanton/tfe-client/tfe/models"
)

type TerraformEnterpriseClient struct {
	client *http.Client
	logger u.ILogger
	host   *url.URL
	token  string

	/*Services*/
	ProviderService *RegistryProviderService
	GpgService      *GpgService
}

func NewClient(logger u.ILogger, cfg *m.ClientConfig) (*TerraformEnterpriseClient, error) {
	hostUrl, err := url.Parse(cfg.Address)
	if err != nil {
		logger.Error("unable to parse the host url.")
		return nil, err
	}
	cli := TerraformEnterpriseClient{
		client: cleanhttp.DefaultPooledClient(),
		logger: logger,
		host:   hostUrl,
		token:  cfg.Token,
	}

	/*Register services*/
	cli.ProviderService = newRegistryProviderService(&cli, logger)
	cli.GpgService = newGpgService(&cli, logger)

	return &cli, nil
}
