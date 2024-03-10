package main

import (
	"github.com/gofiber/fiber/v2"
)

var App *fiber.App

type Module func() []func()

type Controllers struct {
	prefix string
}

func (controllers *Controllers) SetPrefix(prefix string) {
	controllers.prefix = prefix
}

func (controllers *Controllers) Bootstrap(
	baseApp *fiber.App,
	modules ...Module,
) {
	App = fiber.New()
	baseApp.Mount(controllers.prefix, App)

	for _, module := range modules {
		for _, controller := range module() {
			controller()
		}
	}
}

func RegisterControllers() *Controllers {
	controllers := Controllers{prefix: ""}
	return &controllers
}
