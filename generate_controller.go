package main

import "github.com/gofiber/fiber/v2"

func GenerateController(prefix string) *fiber.App {
	fiberApp := fiber.New()
	App.Mount(prefix, fiberApp)
	return fiberApp
}
