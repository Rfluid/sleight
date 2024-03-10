# About

Package to improve module/controller generation in fiber.

# Example

Your fiber application entry point should be something like
```
package main

import (
	user "github.com/Rfluid/go-api/src/module/user/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/Rfluid/fiber-module-controller"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	fiber_module.RegisterControllers().SetPrefix("/api").Bootstrap(
        app,
        user.UserModule,
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