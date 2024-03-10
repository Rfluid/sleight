package fiber_modules

import (
	"github.com/gofiber/fiber/v2"
)

var App *fiber.App

type Module func() []func()

type Modules struct {
	prefix  string
	modules []Module
}

func (modules *Modules) Prefix(
	prefix string,
) *Modules {
	modules.prefix = prefix
	return modules
}

func (modules *Modules) Modules(
	modulesToBootstrap ...Module,
) *Modules {
	modules.modules = modulesToBootstrap
	return modules
}

func (modules *Modules) Bootstrap(
	baseApp *fiber.App,
) {
	App = fiber.New()
	baseApp.Mount(modules.prefix, App)

	for _, module := range modules.modules {
		for _, controller := range module() {
			controller()
		}
	}
}

func RegisterModules() *Modules {
	modules := Modules{prefix: "", modules: []Module{}}
	return &modules
}
