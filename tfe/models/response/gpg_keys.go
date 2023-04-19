package response

type GpgKeys struct {
	Data  []GpgKeyData `json:"data"`
	Links ListLinks
	Meta  ListMeta
}
