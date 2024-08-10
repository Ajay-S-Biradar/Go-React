package routes

import (
	"github.com/Ajay-S-Biradar/Go-React/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupTodoRoutes(app *fiber.App) {
	app.Get("/api/assistant", controllers.GetTask) //.Post("/api/assistant", controllers.AddTask)
	// app.Post("/api/assistant/work", controllers.CreateWork)
}
