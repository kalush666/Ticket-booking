package repositories

import (
	"context"

	"github.com/kalush66/ticket-booking-project-v1/models"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func (r *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {
	events := []*models.Event{}

	if res := r.db.Model(&models.Event{}).Find(&events); res.Error != nil{
		return nil,res.Error
	}
	
	return events,nil
}


func (r *EventRepository) GetOne(ctx context.Context, eventid string) (*models.Event, error) {
	return nil, nil
}

func (r *EventRepository) CreateOne(ctx context.Context, event models.Event) (*models.Event, error) {
	return nil, nil
}

func NewEventRepository(db *gorm.DB) models.EventRepository{
	return &EventRepository{db}
}