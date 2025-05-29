package handlers

import (
	"context"
	"strconv"
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
	eventId,_ := strconv.Atoi(ctx.Params("eventId"))

	context,cancel := context.WithTimeout(context.Background(),time.Duration(5*time.Second))
	defer cancel()

	event, err := h.repository.GetOne(context,uint(eventId))

	if err != nil{
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":"fail",
			"messege":err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":"success",
		"messete":"",
		"data" : event,
	})
}

func (h *EventHandler) CreateOne(ctx *fiber.Ctx) error {
	event := &models.Event{}

	context,cancel := context.WithTimeout(context.Background(),time.Duration(5*time.Second))
	defer cancel()

	if err := ctx.BodyParser(event); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":"fail",
			"messege":err.Error(),
			"data": nil,
		})
	}

	event,err := h.repository.CreateOne(context,event);

	if err != nil{
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":"fail",
			"messege":err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":"success",
		"messete":"event created",
		"data" : event,
	})
}

func (h *EventHandler) UpdateOne(ctx *fiber.Ctx) error {
	eventId, _ := strconv.Atoi(ctx.Params("eventId"))
	updateData := make(map[string]interface{})

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	if err := ctx.BodyParser(&updateData); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  "fail",
				"messege": err.Error(),
				"data":    nil,
		})
	}
	
	event,err := h.repository.UpdateOne(context,uint(eventId),updateData)

	if err != nil{
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":"fail",
			"messege":err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":"success",
		"messete":"event created",
		"data" : event,
	})
}

func (h *EventHandler) DeleteOne(ctx *fiber.Ctx) error{
	eventId, _ := strconv.Atoi(ctx.Params("eventId"))

	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	err := h.repository.DeleteOne(context,uint(eventId))

	if err != nil{
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status": "fail",
			"message": "",
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)

}




func NewEventHandler(router fiber.Router,repository models.EventRepository){
	handler := &EventHandler{
		repository: repository,
	}

	router.Get("/",handler.GetMany)
	router.Get("/", handler.CreateOne)
	router.Get("/:eventId",handler.GetOne)
	router.Put("/:eventId", handler.UpdateOne)
	router.Delete("/:eventId",handler.DeleteOne)
}