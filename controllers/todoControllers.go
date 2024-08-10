package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func GetTask(c *fiber.Ctx) error {
	// Logic to get skills

	return c.Status(200).SendString("hii")
}

// func AddTask(c *fiber.Ctx) error {
// 	// Logic to get skills
// 	l := len()
// 	todo := model.Todo{
// 		ID:l,
// 		body:"do sleep",
// 	}
// 	todos = append(todos,todo)
// 	return c.JSON(todos)
// }
