package middlewares

import (
	"ecommerce-project/utils"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(ctx *fiber.Ctx) error {
	tokenStr := ctx.Cookies("token")

	if tokenStr == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing token",
		})
	}

	claims, err := utils.VerifyToken(tokenStr)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ctx.Locals("user", claims)

	return ctx.Next()
}

func AdminMiddleware(ctx *fiber.Ctx) error {
	claims, err := utils.GetCurrentUser(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}

	if claims.IsAdmin != true {
		return ctx.Status(fiber.StatusForbidden).SendString("Access denied")
	}

	return ctx.Next()
}
