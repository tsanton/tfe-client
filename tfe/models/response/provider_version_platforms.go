package response

type ProviderVersionPlatforms struct {
	Data  []ProviderVersionPlatformData `json:"data"`
	Links ListLinks
	Meta  ListMeta
}
