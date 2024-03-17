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
) Controller {
	c.bootstrap = func() {
		App.Mount(c.prefix, c.App)
		bootstrap()
		defer ControllerBsWg.Done()
	}
	return *c
}

func GenerateController(
	prefix string,
) Controller {
	fiberApp := fiber.New()

	controller := Controller{
		prefix:    prefix,
		bootstrap: func() {},
		App:       fiberApp,
	}

	return controller
}
