package common_response

type PageResponse struct {
	Total int64       `json:"total"`
	List  interface{} `json:"list"`
}
