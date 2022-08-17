package controllers

import (
	"context"
	"net/http"
	"time"

	"Github.com/Guleri24/userapi/configs"
	"Github.com/Guleri24/userapi/models"
	"Github.com/Guleri24/userapi/responses"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func CreateUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel()

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if validatorErr := validate.Struct(&user); validatorErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validatorErr.Error()}})
	}

}
