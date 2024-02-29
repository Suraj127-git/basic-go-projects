package routes

import (
	"os"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/suraj/url-shortener/api/helpers"
	"github.com/suraj/url-shortener/internal/database"
)

type request struct {
	URl         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL            string        `json:"url"`
	CustomShort    string        `json:"short"`
	Expiry         string        `json:"expiry"`
	XRateRemaining int           `json:"rate_limit"`
	XRateLimitRest time.Duration `json:"rate_limit_reached"`
}

const (
	defaultExpiry      = 24 * time.Hour
	apiQuotaTTL        = 30 * time.Minute
	defaultRateLimit   = 10
	rateLimitDecrement = 1
)

func ShortenURL(c *fiber.Ctx) error {
	body := new(request)
	// fmt.Print(body)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	r := database.CreateRedisClient(redisDatabaseMain)
	defer r.Close()

	r2 := database.CreateRedisClient(redisDatabaseIncr)
	defer r2.Close()

	// Rate Limiting
	val, err := r2.Get(c.IP()).Result()
	if err == redis.Nil {
		_ = r2.Set(c.IP(), os.Getenv("API_QUOTA"), apiQuotaTTL).Err()
	} else {
		valInt, _ := strconv.Atoi(val)
		if valInt <= 0 {
			limit, _ := r2.TTL(c.IP()).Result()
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"error":           "Rate limit exceeded",
				"rate_limit_rest": limit / time.Minute,
			})
		}
	}

	// Validate URL
	if !govalidator.IsURL(body.URl) || !helpers.RemoveDomainError(body.URl) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid URL"})
	}

	body.URl = helpers.EnforceHTTP(body.URl)

	// Generate Short URL ID
	var id string
	if body.CustomShort == "" {
		id = uuid.New().String()[:6]
	} else {
		id = body.CustomShort
	}

	// Check if Short URL ID already exists
	val, _ = r.Get(id).Result()
	if val != "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Short URL already exists"})
	}

	// Set Short URL in Redis
	if body.Expiry == 0 {
		body.Expiry = time.Duration(defaultExpiry.Hours()) * time.Hour
	}

	err = r.Set(id, body.URl, time.Duration(body.Expiry)*time.Hour).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unable to connect to server"})
	}

	// Response
	resp := response{
		URL:            body.URl,
		CustomShort:    "",
		Expiry:         time.Duration(body.Expiry).String(),
		XRateRemaining: defaultRateLimit,
		XRateLimitRest: apiQuotaTTL / time.Minute,
	}

	// Update Rate Limiting
	r2.DecrBy(c.IP(), rateLimitDecrement)

	val, _ = r2.Get(c.IP()).Result()
	resp.XRateRemaining, _ = strconv.Atoi(val)

	ttl, _ := r2.TTL(c.IP()).Result()
	resp.XRateLimitRest = ttl / time.Minute

	resp.CustomShort = os.Getenv("DOMAIN") + "/" + id

	return c.Status(fiber.StatusOK).JSON(resp)
}
