# About

Tools to organize your fiber app with module/controller declarations featuring out of the box multithread.

# Example

Your fiber application entry point should be something like
```
package main

import (
	user "github.com/Rfluid/go-api/src/module/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/Rfluid/fiber-modules"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	sleight
        .RegisterModules()
        .Prefix("/api").
        Modules(
            user_module.Main
        ).Bootstrap(
            app,
        )

	app.Listen(":3000")
}
```
Your user.controller.go something like
``` 
package user_controller

func Main() {
    controller := sleight.GenerateController("/user")

    controller.Get("/me", someFunc)
}
```
And your user.module.go
```
package user_module

func Main() {
    GenerateController
}
```