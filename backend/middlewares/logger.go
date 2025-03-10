package middlewares

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Logger() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		start := time.Now()
		err := ctx.Next()
		duration := time.Since(start)
		log.Printf("[%s] %s %s %s", ctx.Method(), ctx.Path(), duration, ctx.IP())
		return err
	}
}