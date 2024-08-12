package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Ajay-S-Biradar/Go-React/config"
	"github.com/Ajay-S-Biradar/Go-React/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Running")

	//removing an item from array using the index of item
	// arr := []int{1,2,3,4,5,6,7,}
	// arr = append(arr[:3], arr[3],arr[5],arr[4])

	// fmt.Println(Todos)
	app := fiber.New()

	err := godotenv.Load(".env")
	PORT := os.Getenv("PORT")

	config.ConnectDB()

	if err != nil {
		log.Fatal("Error loading the .env file", err)
	}

	defer config.DisconnectDB()

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
