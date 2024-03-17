package sleight

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

var ControllerBsWg sync.WaitGroup

type ControllerBootstrap = func()

type Controller struct {
	prefix    string
	bootstrap ControllerBootstrap
	*fiber.App
}

func (c *Controller) SetBootstrap(
	bootstrap ControllerBootstrap,
) {
	c.bootstrap = func() {
		bootstrap()
		defer ControllerBsWg.Done()
	}
}

func GenerateController(
	prefix string,
) Controller {
	fiberApp := fiber.New()
	App.Mount(prefix, fiberApp)

	controller := Controller{
		prefix:    prefix,
		bootstrap: func() {},
		App:       fiberApp,
	}

	return controller
}
