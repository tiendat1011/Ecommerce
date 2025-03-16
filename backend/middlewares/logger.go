package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupLogger() fiber.Handler {

	return logger.New(logger.Config{
		TimeZone:   "Asia/Ho_Chi_Minh",
		TimeFormat: "2002-11-10 19:19:00",
		Output:     os.Stdout,
		Done: func(c *fiber.Ctx, logString []byte) {

		},
	})

	// return func(ctx *fiber.Ctx) error {
	// 	start := time.Now()
	// 	err := ctx.Next()
	// 	duration := time.Since(start)
	// 	log.Printf("[%s] %s %s %s", ctx.Method(), ctx.Path(), duration, ctx.IP())
	// 	return err
	// }
}
