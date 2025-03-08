package handlers

import (
	"ecommerce-project/models"
	"ecommerce-project/services"
	"ecommerce-project/utils"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService *services.UserService
	validator *utils.Validator
}

func NewUserHandler(userService *services.UserService, validator *utils.Validator) *UserHandler {
	return &UserHandler{
		userService: userService,
		validator: validator,
	}
}

func (h *UserHandler) CreateUser(ctx *fiber.Ctx) error {
	u := &models.User{}

	if err := ctx.BodyParser(u); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// Validation
	if errs := h.validator.Validate(u); len(errs) > 0 && errs[0].Error {
		errMsgs := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}

		return fiber.NewError(
			fiber.StatusBadRequest,
			strings.Join(errMsgs, " and "),
		)
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