package handlers

import (
	"ecommerce-project/models"
	"ecommerce-project/services"

	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	categoryService *services.CategoryService
}

func NewCategoryHandler(categoryService *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
	}
}

func (h *CategoryHandler) CreateCategory(ctx *fiber.Ctx) error {
	var category models.Category
	if err := ctx.BodyParser(&category); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	createdCategory, err := h.categoryService.CreateCategory(&category)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(createdCategory)
}