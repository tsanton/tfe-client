package request

type GpgData struct {
	//Type must be "gpg-keys"
	Type       string        `json:"type"`
	Attributes GpgAttributes `json:"attributes"`
}

type GpgAttributes struct {
	AsciiArmor string `json:"ascii-armor"`
	//The namespace of the provider. Must be the same as the organization_name for the provider.
	Namespace string `json:"namespace"`
}

/*##################
### Request body ###
##################*/

type Gpg struct {
	Data GpgData `json:"data"`
}
