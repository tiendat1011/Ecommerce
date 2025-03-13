package handlers

import (
	"ecommerce-project/models"
	"ecommerce-project/services"
	"ecommerce-project/utils"

	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	categoryService services.CategoryService
	validator       *utils.Validator
}

func NewCategoryHandler(categoryService services.CategoryService, validator *utils.Validator) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
		validator:       validator,
	}
}

func (h *CategoryHandler) CreateCategory(ctx *fiber.Ctx) error {
	var c models.Category
	if err := ctx.BodyParser(&c); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if errs := h.validator.Validate(c); len(errs) > 0 && errs[0].Error {
		return h.validator.DefaultMessage(errs)
	}

	createdCategory, err := h.categoryService.CreateCategory(&c)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(createdCategory)
}

func (h *CategoryHandler) UpdateCategory(ctx *fiber.Ctx) error {
	var ur models.UpdateCategoryRequest
	id := ctx.Params("id")

	if err := ctx.BodyParser(&ur); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if errs := h.validator.Validate(&ur); len(errs) > 0 && errs[0].Error {
		return h.validator.DefaultMessage(errs)
	}

	if err := h.categoryService.UpdateCategory(&ur, id); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(ur.Name)
}

func (h *CategoryHandler) DeleteCategory(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := h.categoryService.DeleteCategory(id); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).SendString("Deleted successfully")
}

func (h *CategoryHandler) GetAllCategory(ctx *fiber.Ctx) error {
	category, err := h.categoryService.GetAllCategory()
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(category)
}

func (h *CategoryHandler) GetCategory(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	category, err := h.categoryService.GetCategory(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(category)
}