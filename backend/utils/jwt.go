package utils

import (
	"ecommerce-project/config"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"user_id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(userID string, email string, ctx *fiber.Ctx) error {
	jwtSecret := []byte(config.Cfg.JwtSecret)
	// claims := jwt.MapClaims{
	// 	"username": username,
	// 	"exp":      time.Now().Add(24 * time.Hour).Unix(),
	// }

	claims := Claims{
		UserID: userID,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    jwtToken,
		HTTPOnly: true,
		Secure:   false,
		SameSite: "strict",
		MaxAge:   30 * 24 * 60 * 60,
	})

	return nil
}

func VerifyToken(tokenStr string) (*Claims, error) {
	jwtSecret := []byte(config.Cfg.JwtSecret)
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, errors.New("Invalid token")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("Unauthorized")
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, errors.New("token expired")
	}

	return claims, nil
}

func GetCurrentUser(ctx *fiber.Ctx) (*Claims, error) {
	user, ok := ctx.Locals("user").(*Claims)
	if !ok {
		return nil, errors.New("User not found in context")
	}

	return user, nil
}