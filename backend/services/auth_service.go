package services

import (
	"ecommerce-project/daos"
	"ecommerce-project/models"
	"ecommerce-project/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthService struct {
	userDAO *daos.UserDAO
}

func NewAuthService(userDAO *daos.UserDAO) *AuthService{
	return &AuthService{
		userDAO: userDAO,
	}
}

func (s *AuthService) Login(lr *models.LoginRequest, ctx *fiber.Ctx) error {
	existingUser, err := s.userDAO.GetUserByEmail(lr.Email)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "User not exists")
	}

	if valid := utils.IsPasswordValid(lr.Password, existingUser.Password); !valid {
		return fiber.NewError(fiber.StatusBadRequest, "Password is invalid")
	}

	if err := utils.GenerateToken(existingUser.ID.String(), existingUser.Email, ctx); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

func (s *AuthService) Logout(ctx *fiber.Ctx) error {
	ctx.ClearCookie("token")

	return nil
}