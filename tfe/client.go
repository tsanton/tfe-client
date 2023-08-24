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
	GpgService                     *GpgService
	ProviderService                *RegistryProviderService
	ProviderVersionService         *RegistryProviderVersionService
	ProviderVersionPlatformService *RegistryProviderVersionPlatformService
	SamlConfigurationService       *SamlConfigurationService
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
	cli.GpgService = newGpgService(&cli, logger)
	cli.ProviderService = newRegistryProviderService(&cli, logger)
	cli.ProviderVersionService = newRegistryProviderVersionService(&cli, logger)
	cli.ProviderVersionPlatformService = newRegistryProviderVersionPlatformService(&cli, logger)
	cli.SamlConfigurationService = NewSamlConfigurationService(&cli, logger)

	return &cli, nil
}
