# About

Package to improve module/controller generation in fiber.

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

	fiber_modules.RegisterModules().Prefix("/api").Modules(
        user.UserModule
    ).Bootstrap(
        app,
    )

	app.Listen(":3000")
}
```
Your user.controller.go something like
``` 
func UserController() {
    controller := GenerateController("/user")

    controller.Get("/me", someFunc)
}
```
And your user.module.go
```
func UserModule() {
    return [
        UserController,
    ]
}
```