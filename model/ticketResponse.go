package model

import "svc-myg-ticketing/entity"

type GetDetailTicketResponse struct {
	ListDetailTicket *entity.Ticket      `json:"listDetailTicket" gorm:"foreignKey:Id"`
	ListReplyTicket  []*entity.TicketIsi `json:"listReplyTicket" gorm:"foreignKey:Id"`
}
