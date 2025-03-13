package handlers

import (
	"ecommerce-project/models"
	"ecommerce-project/services"
	"ecommerce-project/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService services.AuthService
	validator *utils.Validator
}

func NewAuthHandler(authService services.AuthService, validator *utils.Validator) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		validator: validator,
	}
}

func (h *AuthHandler) Login(ctx *fiber.Ctx) error {
	lr := &models.LoginRequest{}

	if err := ctx.BodyParser(lr); err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, err.Error())
	}

	// Validation
	if errs := h.validator.Validate(lr); len(errs) > 0 && errs[0].Error {
		return h.validator.DefaultMessage(errs)
	}

	if err := h.authService.Login(lr, ctx); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login success",
	})
}

func (h *AuthHandler) Logout(ctx *fiber.Ctx) error {
	h.authService.Logout(ctx)

	return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "Logout success",
	})
}