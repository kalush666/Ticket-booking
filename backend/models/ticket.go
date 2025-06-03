package models

import (
	"context"
	"time"
)

type Ticket struct {
	ID      uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	EventID uint   `json:"eventId"`
	UserId uint  `json:"userId" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Event   *Event `json:"event,omitempty" gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Entered bool   `json:"entered" default:"false"`
	CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}

type TicketRepository interface {
	GetMany(ctx context.Context,userId uint) ([]*Ticket, error)
	GetOne(ctx context.Context, ticketId uint,userId uint) (*Ticket, error)
	CreateOne(ctx context.Context, ticket *Ticket,userId uint) (*Ticket, error)
	UpdateOne(ctx context.Context, ticketId uint,userId uint, updateData map[string]interface{}) (*Ticket, error)
}

type ValidateTicket struct {
	TicketId uint `json:"ticketId"`
	OwnerId uint `json:"userId"`
}