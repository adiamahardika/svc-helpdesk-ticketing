package model

type GetTerminalRequest struct {
	TerminalId []string `json:"terminalId"`
	Regional   []string `json:"regional"`
	AreaCode   []string `json:"areaCode"`
	GrapariId  []string `json:"grapariId"`
	Status     string   `json:"status"`
}
