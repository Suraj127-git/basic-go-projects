package routes

import (
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/suraj/url-shortener/internal/database"
)

const (
	notFoundStatus      = fiber.StatusNotFound
	internalErrorStatus = fiber.StatusInternalServerError
	redirectStatusCode  = 301
	redisDatabaseMain   = 0
	redisDatabaseIncr   = 1
)

func ResolveURL(c *fiber.Ctx) error {
	url := c.Params("url")

	r := database.CreateRedisClient(redisDatabaseMain)
	defer r.Close()

	value, err := r.Get(url).Result()
	switch {
	case err == redis.Nil:
		return c.Status(notFoundStatus).JSON(fiber.Map{
			"error": "Short not found in the database",
		})
	case err != nil:
		return c.Status(internalErrorStatus).JSON(fiber.Map{
			"error": "Cannot connect to the database",
		})
	}

	_ = r.Incr("counter")

	return c.Redirect(value, redirectStatusCode)
}
