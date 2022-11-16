package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	// "github.com/hasirm/goauth/pkg/controllers"
	"github.com/hasirm/goauth/pkg/database"
	"github.com/hasirm/goauth/pkg/routes"
)

func main() {
	fmt.Println("server running....")

	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	// controllers.HandleRequests()

	app.Listen(":8000")

}
