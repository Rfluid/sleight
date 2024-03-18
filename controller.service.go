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
	fiber.Router
}

func (c *Controller) setRouter() {
	c.Router = App.Group(c.prefix)
}

func (c *Controller) SetBootstrap(
	bootstrap ControllerBootstrap,
) Controller {
	c.bootstrap = func() {
		c.setRouter()
		bootstrap()
		defer ControllerBsWg.Done()
	}
	return *c
}

func GenerateController(
	prefix string,
) Controller {
	controller := Controller{
		prefix:    prefix,
		bootstrap: func() {},
	}

	return controller
}
