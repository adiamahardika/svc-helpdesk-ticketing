package model

type StandardResponse struct {
	HttpStatusCode int      `json:"http_status_code"`
	ResponseCode   string   `json:"response_code"`
	Description    []string `json:"description"`
}
