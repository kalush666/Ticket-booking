package repositories

import (
	"context"

	"github.com/kalush66/ticket-booking-project-v1/models"
	"gorm.io/gorm"
)

type TicketRepository struct {
	db *gorm.DB
}

func (r *TicketRepository) GetMany(ctx context.Context,userId uint) ([]*models.Ticket, error) {
	tickets := []*models.Ticket{}

	res := r.db.Model(&models.Ticket{}).Where("user_id = ?",userId).Preload("Event").Order("updated_at DESC").Find(&tickets)

	if res.Error != nil {
		return nil, res.Error
	}

	return tickets, nil
}

func (r *TicketRepository) GetOne(ctx context.Context, ticketId uint,userId uint) (*models.Ticket, error) {
	ticket := &models.Ticket{}

	res := r.db.Model(ticket).Where("id = ?",ticketId).Where("user_id = ?",userId).Preload("Event").Where("id = ?", ticketId).First(ticket)

	if res.Error != nil {
		return nil, res.Error
	}

	return ticket, nil
}

func (r *TicketRepository) CreateOne(ctx context.Context, ticket *models.Ticket,userId uint) (*models.Ticket, error) {
	ticket.UserId = userId

	res := r.db.Model(&models.Ticket{}).Create(ticket)

	if res.Error != nil {
		return nil, res.Error
	}

	return r.GetOne(ctx,ticket.ID,userId)
}

func (r *TicketRepository) UpdateOne(ctx context.Context, ticketId uint,userId uint, updateData map[string]interface{}) (*models.Ticket, error) {
	ticket := &models.Ticket{}

	res := r.db.Model(ticket).Where("id = ?", ticketId).Updates(updateData)

	if res.Error != nil {
		return nil, res.Error
	}

	return r.GetOne(ctx, ticketId,userId)
}

func NewTicketRepository(db *gorm.DB) models.TicketRepository {
	return &TicketRepository{
		db,
	}
}