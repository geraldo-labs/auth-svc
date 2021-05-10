package main

import (
	"github.com/geraldo-labs/auth-svc/internal/config"
	"github.com/geraldo-labs/auth-svc/pkg/database"
	"github.com/geraldo-labs/auth-svc/pkg/logging"
	"github.com/geraldo-labs/auth-svc/pkg/tokens"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	log := logging.New()
	conf := config.New("", log)
	app := fiber.New()

	db := database.New(conf.DatabaseDSN, log)
	tokenSvc := tokens.New(db)

	app.Get("*", func(c *fiber.Ctx) error {
		authorization := c.Get("Authorization")
		header := tokens.Header{}
		header.Set("Authorization", authorization)
		if err := tokenSvc.Validate(header); err != nil {
			log.Error(err.Error())
			return c.SendStatus(http.StatusUnauthorized)
		}

		return c.SendStatus(http.StatusNoContent)
	})

	// Respond all Options as OK
	app.Options("*", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusNoContent)
	})

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
