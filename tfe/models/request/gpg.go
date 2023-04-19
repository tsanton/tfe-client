package request

type Gpg struct {
	Data GpgData `json:"data"`
}

type GpgData struct {
	//Type must be "gpg-keys"
	Type       string            `json:"type"`
	Attributes GpgDataAttributes `json:"attributes"`
}

type GpgDataAttributes struct {
	AsciiArmor string `json:"ascii-armor"`
	//The namespace of the provider. Must be the same as the organization_name for the provider.
	Namespace string `json:"namespace"`
}
