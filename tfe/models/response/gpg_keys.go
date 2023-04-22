package response

type GpgKeys struct {
	Data  []gpgKeyData `json:"data"`
	Links listLinks
	Meta  listMeta
}
