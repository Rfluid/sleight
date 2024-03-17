package sleight

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

type Modules struct {
	prefix  string
	modules []Module
}

func Register() *Modules {
	modules := Modules{prefix: "", modules: []Module{}}
	return &modules
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

	var modulesWg sync.WaitGroup

	for _, module := range modules.modules {
		modulesWg.Add(1)

		go func(module Module) {
			module.Bootstrap()

			defer modulesWg.Done()
		}(module)
	}

	modulesWg.Wait()
}
