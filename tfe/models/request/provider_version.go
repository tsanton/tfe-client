package request

type ProviderVersion struct {
	Data ProviderVersionData `json:"data"`
}

type ProviderVersionData struct {
	//Type must be 'registry-provider-versions'
	Type       string                        `json:"type"`
	Attributes ProviderVersionDataAttributes `json:"attributes"`
}

type ProviderVersionDataAttributes struct {
	Version   string   `json:"version"`
	KeyId     string   `json:"key-id"`
	Protocols []string `json:"protocols"`
}
