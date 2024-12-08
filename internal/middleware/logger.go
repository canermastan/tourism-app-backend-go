package middleware

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

func LoggerMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		log.Printf("Incoming Request: %s %s", c.Method(), c.OriginalURL())

		err := c.Next()

		log.Printf("Completed Request: %s %s, Duration: %s",
			c.Method(),
			c.OriginalURL(),
			time.Since(start),
		)

		return err
	}
}
