package main

import (
    "fiber-mongo-api/config"
    "fiber-mongo-api/route" 
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    //run database
    config.ConnectDB()

    //routes
    route.UserRoute(app) 

    app.Listen(":7000")
}