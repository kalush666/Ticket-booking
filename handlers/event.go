package handlers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kalush66/ticket-booking-project-v1/models"
)

type EventHandler struct{
	repository models.EventRepository
}

func (h *EventHandler) GetMany(ctx *fiber.Ctx) error {
	context,cancel := context.WithTimeout(context.Background(),time.Duration(5*time.Second))
	defer cancel()

	events,error := h.repository.GetMany(context)

	if error != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status" : "fail",
			"message": "Failed to fetch events",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"message": "Events fetched successfully",
		"data": events,
	})
}

func (h *EventHandler) GetOne(ctx *fiber.Ctx) error {
	return nil
}

func (h *EventHandler) CreateOne(ctx *fiber.Ctx) error {
	return nil
}



func NewEventHandler(router fiber.Router,repository models.EventRepository){
	handler := &EventHandler{
		repository: repository,
	}

	router.Get("/",handler.GetMany)
	router.Get("/", handler.CreateOne)
	router.Get("/:eventId",handler.GetOne)
}