package services

import (
	"ecommerce-project/daos"
	"ecommerce-project/models"

	"github.com/gofiber/fiber/v2"
)

type CategoryService interface {
	CreateCategory(category *models.Category) (*models.Category, error)
	UpdateCategory(ur *models.UpdateCategoryRequest, id string) error
	DeleteCategory(id string) error
	GetAllCategory() ([]*models.Category, error)
	GetCategory(id string) (*models.Category, error)
}

type categoryService struct {
	categoryDAO daos.CategoryDAO
}

func NewCategoryService(categoryDAO daos.CategoryDAO) *categoryService {
	return &categoryService{
		categoryDAO: categoryDAO,
	}
}

func (s *categoryService) CreateCategory(category *models.Category) (*models.Category, error) {
	if existingCategory, _ := s.categoryDAO.GetCategoryByName(category.Name); existingCategory != nil {
		return nil, fiber.NewError(fiber.StatusConflict, "Category already exists")
	}

	createdCategory, err := s.categoryDAO.CreateCategory(category)
	if err != nil {
		return nil, err
	}

	return createdCategory, nil
}

func (s *categoryService) UpdateCategory(ur *models.UpdateCategoryRequest, id string) error {
	if err := s.categoryDAO.UpdateCategory(ur, id); err != nil {
		return err
	}

	return nil
}

func (s *categoryService) DeleteCategory(id string) error {
	if err := s.categoryDAO.DeleteCategory(id); err != nil {
		return err
	}

	return nil
}

func (s *categoryService) GetAllCategory() ([]*models.Category, error) {
	category, err := s.categoryDAO.GetAllCategory()
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *categoryService) GetCategory(id string) (*models.Category, error) {
	category, err := s.categoryDAO.GetCategory(id)
	if err != nil {
		return nil, err
	}

	return category, nil
}