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

func GenerateController(
	prefix string,
	bootstrap func(),
) Controller {
	fiberApp := fiber.New()
	App.Mount(prefix, fiberApp)

	controller := Controller{
		prefix: prefix,
		bootstrap: func() {
			bootstrap()
			defer ControllerBsWg.Done()
		},
		App: fiberApp,
	}

	return controller
}
