package entity

type User struct {
	Id         int    `json:"id" gorm:"primaryKey"`
	Username   string `json:"username"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Status     string `json:"status"`
	AreaCode   string `json:"areaCode"`
	Regional   string `json:"regional"`
	GrapariId  string `json:"grapariId"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
	Roles      string `json:"roles"`
	TerminalId string `json:"terminalId"`
	RuleId     int    `json:"ruleId"`
}
