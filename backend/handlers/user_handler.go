package handlers

import (
	"ecommerce-project/models"
	"ecommerce-project/services"
	"ecommerce-project/utils"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService services.UserService
	validator   *utils.Validator
}

func NewUserHandler(userService services.UserService, validator *utils.Validator) *UserHandler {
	return &UserHandler{
		userService: userService,
		validator:   validator,
	}
}

func (h *UserHandler) CreateUser(ctx *fiber.Ctx) error {
	u := &models.User{}

	if err := ctx.BodyParser(u); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// Validation
	if errs := h.validator.Validate(u); len(errs) > 0 && errs[0].Error {
		return h.validator.DefaultMessage(errs)
	}

	createdUser, err := h.userService.CreateUser(u)
	if err != nil {
		return err
	}

	// Generate token and set cookie
	if err := utils.GenerateToken(createdUser.ID.Hex(), createdUser.Email, createdUser.IsAdmin, ctx); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(createdUser)
}

func (h *UserHandler) GetAllUsers(ctx *fiber.Ctx) error {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(users)
}

func (h *UserHandler) GetUserProfile(ctx *fiber.Ctx) error {
	user, err := h.userService.GetUserProfile(ctx)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}

func (h *UserHandler) UpdateUserProfile(ctx *fiber.Ctx) error {
	ur := &models.UpdateRequest{}
	if err := ctx.BodyParser(ur); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if err := h.userService.UpdateUserProfile(ur, ctx); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON("Updated successfully")
}

func (h *UserHandler) GetUserById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, err := h.userService.GetUserById(id)
	if err != nil {
		return err 
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}

func (h *UserHandler) UpdateUserById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	ur := &models.UpdateRequest{}
	if err := ctx.BodyParser(ur); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if err := h.userService.UpdateUserById(ur, id); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON("")
}

func (h *UserHandler) DeleteUserById(ctx *fiber.Ctx) error {
	h.userService.DeleteUserById(ctx.Params("id"))

	return ctx.Status(fiber.StatusOK).SendString("Deleted successfully")
}