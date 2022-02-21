package model

type StandardResponse struct {
	HttpStatus  int      `json:"http_status"`
	StatusCode  string   `json:"status_code"`
	Description []string `json:"description"`
}
