package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kalush66/ticket-booking-project-v1/config"
	"github.com/kalush66/ticket-booking-project-v1/db"
	"github.com/kalush66/ticket-booking-project-v1/handlers"
	"github.com/kalush66/ticket-booking-project-v1/repositories"
	"github.com/kalush66/ticket-booking-project-v1/services"
)

func main() {
	envconfig := config.NewEnvConfig()
	db := db.Init(envconfig,db.DBMigrator)

	 app := fiber.New(fiber.Config{
		AppName: "Ticket Booking API",
		ServerHeader: "Fiber",
	}) 

	eventRepository := repositories.NewEventRepository(db)
	ticketRepository := repositories.NewTicketRepository(db)
	authRepository := repositories.NewAuthRepository(db)

	authService := services.NewAuthService(authRepository)

	server := app.Group("/api")
	handlers.NewAuthHandler(server.Group("/auth"), authService)

	privateRoutes := server.Use(middleware.AuthProtected(db))

	handlers.NewEventHandler(server.Group("/event"),eventRepository)
	handlers.NewTicketHandler(server.Group("/ticket"),ticketRepository)

	if err := app.Listen(fmt.Sprintf(":%s", envconfig.ServerPort)); err != nil {
		log.Fatal(err)
	}
}