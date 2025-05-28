package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kalush66/ticket-booking-project-v1/config"
	"github.com/kalush66/ticket-booking-project-v1/db"
	"github.com/kalush66/ticket-booking-project-v1/handlers"
	"github.com/kalush66/ticket-booking-project-v1/repositories"
)

func main() {
	envconfig := config.NewEnvConfig()
	db := db.Init(envconfig,db.DBMigrator)

	 app := fiber.New(fiber.Config{
		AppName: "Ticket Booking API",
		ServerHeader: "Fiber",
	}) 

	eventRepository := repositories.NewEventRepository(db)

	server := app.Group("/api")

	handlers.NewEventHandler(server.Group("/event"),eventRepository)

	if err := app.Listen(":3030"); err != nil {
        log.Fatal(err)
    }
}