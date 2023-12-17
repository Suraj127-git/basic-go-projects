package routes

import (
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/suraj/url-shortener/internal/database"
)

func ResolveURL(c *fiber.Ctx) error {
	url := c.Params("url")

	r := database.CreateRedisClient(0)
	defer r.Close()

	value, err := r.Get(url).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Short not found in the database",
		})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot connect to database",
		})
	}

	r_inr := database.CreateRedisClient(1)
	defer r_inr.Close()

	_ = r_inr.Incr("counter")

	return c.Redirect(value, 301)
}
