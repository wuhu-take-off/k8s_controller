package common_request

type PageRequest struct {
	PageNum  int `json:"pageNum"`
	PageSize int `json:"pageSize"`
}
