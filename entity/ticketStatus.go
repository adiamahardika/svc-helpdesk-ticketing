package entity

type TicketStatus struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Index     int    `json:"index"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
