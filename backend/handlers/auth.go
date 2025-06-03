package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/kalush66/ticket-booking-project-v1/models"
)

var validate = validator.New()

type AuthHandler struct{
	service models.AuthService
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	creds := &models.AuthCredentials{}

	context,cancel := context.WithTimeout(context.Background(),time.Duration(5*time.Second))
	defer cancel()

	if err := c.BodyParser(&creds); err !=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":"fail",
			"message":err.Error(),
			"data":nil,
		})
	}

	if err := validate.Struct(creds); err !=nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":"fail",
			"message":err.Error(),
			"data":nil,
		})
	}

	token,user,err := h.service.Login(context,creds)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":"fail",
			"message":err.Error(),
			"data":nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":"success",
		"message":"Login successful",
		"data": &fiber.Map{
			"token": token,
			"user":  user,
		},
	})
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	creds := &models.AuthCredentials{}

	context,cancel := context.WithTimeout(context.Background(),time.Duration(5*time.Second))
	defer cancel()

	if err := c.BodyParser(&creds); err !=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":"fail",
			"message":err.Error(),
			"data":nil,
		})
	}

	if err := validate.Struct(creds); err !=nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":"fail",
			"message":fmt.Errorf("please provide valid credentials").Error(),
		})
	}

	token,user,err := h.service.Register(context,creds)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":"fail",
			"message":err.Error(),
			"data":nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":"success",
		"message":"Login successful",
		"data": &fiber.Map{
			"token": token,
			"user":  user,
		},
	})
}

func NewAuthHandler(router fiber.Router,service models.AuthService){
	handler := &AuthHandler{
		service,
	}

	router.Post("/login", handler.Login)
	router.Post("/register", handler.Register)
}