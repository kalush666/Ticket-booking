package repositories

import (
	"context"

	"github.com/kalush66/ticket-booking-project-v1/models"
	"time"
)

type EventRepository struct {
	db any
}

func (r *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {
	events := []*models.Event{}

	events = append(events, &models.Event{
		ID:       "1",
		Name:     "Concert",
		Location: "Stadium",
		Date:     time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	return events, nil
}


func (r *EventRepository) GetOne(ctx context.Context, eventid string) (*models.Event, error) {
	return nil, nil
}

func (r *EventRepository) CreateOne(ctx context.Context, event models.Event) (*models.Event, error) {
	return nil, nil
}

func NewEventRepository(db any) models.EventRepository{
	return &EventRepository{db}
}