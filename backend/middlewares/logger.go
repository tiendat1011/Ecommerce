package middlewares

import (
	"bytes"
	"ecommerce-project/config"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupLogger() fiber.Handler {
	file, err := os.OpenFile("./logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error opening log file: %v",err)
		return nil
	}

	multiWriter := io.MultiWriter(os.Stdout, file)

	return logger.New(logger.Config{
		TimeZone:   "Asia/Ho_Chi_Minh",
		TimeFormat: "2002-11-10 19:19:00",
		Output:     multiWriter,
		Done: func(c *fiber.Ctx, logString []byte) {
			SendLogToDiscord(string(logString))
		},
	})
}

type DiscordPayload struct {
	Content string `json:"content"`
}

func SendLogToDiscord(message string) {
	payload := DiscordPayload{Content: "**API Log** ```" + message + "```"}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Fatal("Unable to marshal: %v", err.Error())
	}

	http.Post(config.Cfg.DiscordWebhookURL, "application/json", bytes.NewBuffer(jsonPayload))
}
