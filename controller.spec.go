package fiber_modules

import "github.com/gofiber/fiber/v2"

type Controller = *fiber.App

func GenerateController(prefix string) Controller {
	fiberApp := fiber.New()
	App.Mount(prefix, fiberApp)
	return fiberApp
}
