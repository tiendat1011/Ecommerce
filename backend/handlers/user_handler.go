package handlers

import (
	"ecommerce-project/models"
	"ecommerce-project/services"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) CreateUser(ctx *fiber.Ctx) error {
	u := &models.User{}

	if err := ctx.BodyParser(u); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	createdUser, err := h.userService.CreateUser(u)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusCreated).JSON(createdUser)
}

func (h *UserHandler) GetUser(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).SendString("Hello World")
}