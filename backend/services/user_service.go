package services

import (
	"ecommerce-project/daos"
	"ecommerce-project/models"
	"ecommerce-project/utils"

	"github.com/gofiber/fiber/v2"
)

type UserService struct {
	userDAO *daos.UserDAO
}

func NewUserService(uDAO *daos.UserDAO) *UserService {
	return &UserService{
		userDAO: uDAO,
	}
}

func (s *UserService) CreateUser(user *models.User) (*models.User, error) {
	if exists, _ := s.userDAO.GetUserByEmail(user.Email); exists != nil {
		return nil, fiber.NewError(fiber.StatusConflict, "User already exists")
	}

	hashedPassword, err := utils.HashPassword(user.Username)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	user.Password = hashedPassword

	createdUser, err := s.userDAO.CreateUser(user)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return createdUser, nil
}

func (s *UserService) GetAllUsers() ([]*models.User, error) {
	return s.userDAO.GetAllUsers()
}

func (s *UserService) GetUserProfile(ctx *fiber.Ctx) (*models.User, error) {
	userClaims, err := utils.GetCurrentUser(ctx)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	user, err := s.userDAO.GetUserById(userClaims.UserID)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return user, nil
}
