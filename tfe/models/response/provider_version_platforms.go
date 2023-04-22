package response

type ProviderVersionPlatforms struct {
	Data  []providerVersionPlatformData `json:"data"`
	Links listLinks
	Meta  listMeta
}
