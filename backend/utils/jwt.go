package utils

import (
	"ecommerce-project/config"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(username string, ctx *fiber.Ctx) error {
	jwtSecret := []byte(config.Cfg.JwtSecret)
	claims := jwt.MapClaims{
		"username": username,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	ctx.Cookie(&fiber.Cookie{
		Name: "token",
		Value: jwtToken,
		HTTPOnly: true,
		Secure: false,
		SameSite: "strict",
		MaxAge: 30 * 24 * 60 * 60 * 1000,
	})

	return nil
}