package services

import (
	"ecommerce-project/daos"
	"ecommerce-project/models"
	"ecommerce-project/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthService interface {
	Login(lr *models.LoginRequest, ctx *fiber.Ctx) error
	Logout(ctx *fiber.Ctx) error
}

type authService struct {
	userDAO daos.UserDAO
}

func NewAuthService(userDAO daos.UserDAO) *authService{
	return &authService{
		userDAO: userDAO,
	}
}

func (s *authService) Login(lr *models.LoginRequest, ctx *fiber.Ctx) error {
	existingUser, err := s.userDAO.GetUserByEmail(lr.Email)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "User not exists")
	}

	if valid := utils.IsPasswordValid(lr.Password, existingUser.Password); !valid {
		return fiber.NewError(fiber.StatusBadRequest, "Password is invalid")
	}

	if err := utils.GenerateToken(existingUser.ID.Hex(), existingUser.Email, existingUser.IsAdmin, ctx); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

func (s *authService) Logout(ctx *fiber.Ctx) error {
	ctx.ClearCookie("token")

	return nil
}