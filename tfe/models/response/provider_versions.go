package response

type ProviderVersions struct {
	Data  []providerVersionData `json:"data"`
	Links listLinks
	Meta  listMeta
}
