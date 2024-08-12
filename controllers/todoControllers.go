package controllers

import (
	"context"
	"fmt"
	"log"

	store "github.com/Ajay-S-Biradar/Go-React/Store"
	"github.com/Ajay-S-Biradar/Go-React/config"
	"github.com/Ajay-S-Biradar/Go-React/model"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTask(c *fiber.Ctx) error {
	// Logic to get skills
	cursor, err := config.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal("Error in getTask", err)
	}

	store.Todos = []model.Todo{}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo model.Todo
		if err := cursor.Decode(&todo); err != nil {
			return err
		}

		store.Todos = append(store.Todos, todo)
	}
	fmt.Println(store.Todos)
	return c.Status(200).JSON(store.Todos)
}

func AddTask(c *fiber.Ctx) error {
	// config.Collection.InsertOne()
	todo := new(model.Todo)
	if err := c.BodyParser(todo); err != nil {
		return err
	}

	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Body is empty"})
	}

	insertedResult, err := config.Collection.InsertOne(context.Background(), todo)
	if err != nil {
		return err
	}

	todo.ID = insertedResult.InsertedID.(primitive.ObjectID)

	return c.JSON(fiber.Map{"success": true, "addedTodo": todo})
}

func CheckTask(c *fiber.Ctx) error {
	// c.BodyParser()
	id := c.Params("id")
	fmt.Println("id:", id)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil
	}
	fmt.Println("objID:", objectID)
	// err = config.Collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&todo)
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"completed": true}}
	_, err = config.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"success": true})
}

func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.JSON(fiber.Map{"Error": "Error while converting"})
	}
	_, err = config.Collection.DeleteOne(context.Background(), bson.M{"_id": objectId})
	if err != nil {
		return c.JSON(fiber.Map{"Error": err})
	}

	return c.JSON(fiber.Map{"success": true})
}
