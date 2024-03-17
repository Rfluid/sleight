package fiber_modules

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

type ModuleBootstrap func()

type Modules struct {
	prefix  string
	modules []ModuleBootstrap
}

func (modules *Modules) Prefix(
	prefix string,
) *Modules {
	modules.prefix = prefix
	return modules
}

func (modules *Modules) Modules(
	modulesToBootstrap ...ModuleBootstrap,
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

		go func(module ModuleBootstrap) {
			module()

			defer modulesWg.Done()
		}(module)
	}

	modulesWg.Wait()
}

func RegisterModules() *Modules {
	modules := Modules{prefix: "", modules: []ModuleBootstrap{}}
	return &modules
}
