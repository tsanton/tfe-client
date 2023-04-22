package response

type ProviderVersions struct {
	Data  []ProviderVersionData `json:"data"`
	Links ListLinks
	Meta  ListMeta
}
