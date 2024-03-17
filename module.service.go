package sleight

type ModuleBootstrap func()

type Module struct {
	controllers []Controller
}

func (module *Module) SetControllers(
	controllers ...Controller,
) *Module {
	module.controllers = controllers
	return module
}

func (module *Module) Bootstrap() *Module {
	for _, controller := range module.controllers {
		ControllerBsWg.Add(1)

		go controller.bootstrap()
	}

	ControllerBsWg.Wait()
	return module
}
