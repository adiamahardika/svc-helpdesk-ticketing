package model

type StandardResponse struct {
	HttpStatusCode int      `json:"httpStatusCode"`
	ResponseCode   string   `json:"responseCode"`
	Description    []string `json:"description"`
}
