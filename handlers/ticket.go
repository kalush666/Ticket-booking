package handlers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kalush66/ticket-booking-project-v1/models"
)

type TicketHandler struct {
	repository models.TicketRepository
}


func (h *TicketHandler) GetMany(ctx *fiber.Ctx) error {
	context,cancel := context.WithTimeout(context.Background(),time.Duration(5*time.Second))
	defer cancel()

	ticket,err := h.repository.GetMany(context)

	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status" : "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"message": "Tickets fetched successfully",
		"data": ticket,
	})
}

func (h *TicketHandler) GetOne(ctx *fiber.Ctx) error {
	context,cancel := context.WithTimeout(context.Background(),time.Duration(5*time.Second))
	defer cancel()

	ticketId,_ := strconv.Atoi(ctx.Params("ticketId"))
	ticket,err := h.repository.GetOne(context,uint(ticketId))

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"status" : "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"message": "Ticket fetched successfully",
		"data": ticket,
	})
}

func (h *TicketHandler) CreateOne(ctx *fiber.Ctx) error {
	context,cancel := context.WithTimeout(context.Background(),time.Duration(5*time.Second))
	defer cancel()

	ticket := &models.Ticket{}

	if err := ctx.BodyParser(ticket); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status" : "fail",
			"message": err.Error(),
		})
	}

	ticketRes,err := h.repository.CreateOne(context,ticket)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status" : "fail",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status": "success",
		"message": "Ticket created successfully",
		"data": ticketRes,
	})
}

func (h *TicketHandler) ValidateOne(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	validateBody := &models.ValidateTicket{}

	if ctx.Get("Content-Type") != "application/json" {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Content-Type must be application/json",
			"data":    nil,
		})
	}

	if err := ctx.BodyParser(validateBody); err != nil {
		fmt.Printf("BodyParser error: %v\n", err)
		fmt.Printf("Raw body: %s\n", string(ctx.Body()))
		
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status":  "fail",
			"message": fmt.Sprintf("Failed to parse request body: %v", err),
			"data":    nil,
		})
	}

	if validateBody.TicketId <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "ticketId is required and must be greater than 0",
			"data":    nil,
		})
	}

	validateData := make(map[string]interface{})
	validateData["entered"] = true

	ticket, err := h.repository.UpdateOne(context, uint(validateBody.TicketId), validateData)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Ticket validated successfully",
		"data":    ticket,
	})
}

func NewTicketHandler(router fiber.Router, repository models.TicketRepository){
	handler := &TicketHandler{
		repository,
	}

	router.Get("/", handler.GetMany)
	router.Get("/:ticketId", handler.GetOne)
	router.Post("/", handler.CreateOne)
	router.Post("/validate", handler.ValidateOne)
}