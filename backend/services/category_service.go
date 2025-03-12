package services

import (
	"ecommerce-project/daos"
	"ecommerce-project/models"

	"github.com/gofiber/fiber/v2"
)

type CategoryService struct {
	categoryDAO *daos.CategoryDAO
}

func NewCategoryService(categoryDAO *daos.CategoryDAO) *CategoryService {
	return &CategoryService{
		categoryDAO: categoryDAO,
	}
}

func (s *CategoryService) CreateCategory(category *models.Category) (*models.Category, error) {
	if existingCategory, _ := s.categoryDAO.GetCategoryByName(category.Name); existingCategory != nil {
		return nil, fiber.NewError(fiber.StatusConflict, "Category already exists")
	}

	createdCategory, err := s.categoryDAO.CreateCategory(category)
	if err != nil {
		return nil, err
	}

	return createdCategory, nil
}

func (s *CategoryService) UpdateCategory(ur *models.UpdateCategoryRequest, id string) error {
	if err := s.categoryDAO.UpdateCategory(ur, id); err != nil {
		return err
	}

	return nil
}

func (s *CategoryService) DeleteCategory(id string) error {
	if err := s.categoryDAO.DeleteCategory(id); err != nil {
		return err
	}

	return nil
}

func (s *CategoryService) GetAllCategory() ([]*models.Category, error) {
	category, err := s.categoryDAO.GetAllCategory()
	if err != nil {
		return nil, err
	}

	return category, nil
}