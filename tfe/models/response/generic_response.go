package response

type listLinks struct {
	Self  string `json:"self"`
	First string `json:"first"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
	Last  string `json:"last"`
}

type listMeta struct {
	Pagination listMetaPagination
}

type listMetaPagination struct {
	CurrentPage int  `json:"current-page"`
	PageSize    int  `json:"page-size"`
	PrevPage    *int `json:"prev-page"`
	NextPage    *int `json:"next-page"`
	TotalPages  int  `json:"total-pages"`
	TotalCount  int  `json:"total-count"`
}
