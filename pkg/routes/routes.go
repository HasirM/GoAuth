package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/hasirm/vaddiapp/auth/pkg/controllers"
)

func Setup(app *fiber.App) {
	//APIs
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Post("/api/logout", controllers.Logout)
	app.Post("/api/refresh", controllers.Refresh)
	app.Post("/api/validate", controllers.Validate)


	// // Unauthenticated route
	// app.Get("/home", controllers.Home)
	// app.Get("/accessible", controllers.Accessible)

	// JWT Middleware
	// check whether the ajwt is valid or not, if ajwt is expired and rjwt is still valid then
	// creates a new pair of JWT's.
	app.Use(controllers.Refresh, jwtware.New(jwtware.Config{
		SigningKey:  []byte(controllers.SecretKey),
		TokenLookup: "cookie:ajwt",
	}))

	// // Restricted Routes
	// app.Get("/restricted", controllers.Restricted)
	// app.Get("/home1", controllers.Home1)
	// app.Get("/home2", controllers.Home2)
	// app.Get("/home3", controllers.Home3)
}
