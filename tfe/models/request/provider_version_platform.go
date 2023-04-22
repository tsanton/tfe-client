package request

type ProviderVersionPlatform struct {
	Data ProviderVersionPlatformData `json:"data"`
}

type ProviderVersionPlatformData struct {
	//Type must be 'registry-provider-version-platforms'
	Type       string                                `json:"type"`
	Attributes ProviderVersionPlatformDataAttributes `json:"attributes"`
}

type ProviderVersionPlatformDataAttributes struct {
	Os       string `json:"os"`
	Arch     string `json:"arch"`
	Shasum   string `json:"shasum"`
	Filename string `json:"filename"`
}
