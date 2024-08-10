package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Ajay-S-Biradar/Go-React/model"
	"github.com/Ajay-S-Biradar/Go-React/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var todos []model.Todo

func main() {
	fmt.Println("Running")
	fmt.Println(todos)
	app := fiber.New()

	err := godotenv.Load(".env")
	PORT := os.Getenv("PORT")

	if err != nil {
		log.Fatal("Error loading the .env file", err)
	}

	routes.SetupTodoRoutes(app)

	// app.Get("/", func(c *fiber.Ctx) error {
	// info := map[string]string{
	// 	"name": "ajay",
	// 	"age":  "20",
	// }
	// return c.Status(200).SendString("Working...")
	//SendString("hii") - can do this also but need to send the json data right
	// })

	// app.Get("/api/assistant", func(c *fiber.Ctx) error {
	// 	return c.Status(200).JSON(todo)
	// })

	// app.Post("/api/assistant", func(c *fiber.Ctx) error{
	// 	t:= Todo{}
	// return c.Status(200).
	// })

	log.Fatal(app.Listen(":" + PORT))
}
