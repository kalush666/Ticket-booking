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

	if res := r.db.Model(&models.Event{}).Order("updated_at desc").Find(&events); res.Error != nil{
		return nil,res.Error
	}
	return events,nil
}


func (r *EventRepository) GetOne(ctx context.Context, eventid uint) (*models.Event, error) {
	event := &models.Event{}

	res := r.db.Model(event).Where("id = ?",eventid).First(event)

	if res.Error != nil {
		return nil,res.Error
	}

	return event,nil
}

func (r *EventRepository) CreateOne(ctx context.Context, event *models.Event) (*models.Event, error) {
	res := r.db.Model(event).Create(event)

	if res.Error != nil {
		return nil,res.Error
	}

	return event,nil
}

func (r *EventRepository) UpdateOne(ctx context.Context, eventId uint,updateData map[string]interface{}) (*models.Event, error) {
	event := &models.Event{}

	updateRes := r.db.Model(event).Where("id = ?", eventId).Updates(updateData)

	if updateRes.Error != nil {
		return nil, updateRes.Error
	}

	getResponse := r.db.Model(event).Where("id = ?", eventId).First(event)

	if getResponse.Error != nil {
		return nil, getResponse.Error
	}

	return event, nil
}

func (r *EventRepository) DeleteOne(ctx context.Context, eventId uint) error {
	res := r.db.Model(&models.Event{}).Where("id = ?", eventId).Delete(&models.Event{})
	return res.Error
}

func NewEventRepository(db *gorm.DB) models.EventRepository{
	return &EventRepository{db}
}