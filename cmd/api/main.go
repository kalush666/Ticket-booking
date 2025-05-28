package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/kalush66/ticket-booking-project-v1/env"
	"github.com/kalush66/ticket-booking-project-v1/handlers"
	"github.com/kalush66/ticket-booking-project-v1/repositories"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Print("Error loading .env file")
	}else{
		log.Print(".env loaded")
	}
	

	port := env.GetString("PORT", ":3000")


	 app := fiber.New(fiber.Config{
		AppName: "Ticket Booking API",
		ServerHeader: "Fiber",
	}) 

	eventRepository := repositories.NewEventRepository(nil)

	server := app.Group("/api")

	handlers.NewEventHandler(server.Group("/event"),eventRepository)

	if err := app.Listen(port); err != nil {
        log.Fatal(err)
    }
}