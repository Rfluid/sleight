package sleight

type ModuleBootstrap func()

type Module struct {
	Controllers []Controller
}

func (module *Module) SetControllers(
	controllers ...Controller,
) *Module {
	module.Controllers = controllers
	return module
}

func (module *Module) Bootstrap() *Module {
	for _, controller := range module.Controllers {
		ControllerBsWg.Add(1)

		go controller.bootstrap()
	}

	ControllerBsWg.Wait()
	return module
}
