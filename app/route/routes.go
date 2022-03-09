package route

import (
    "fiber-mongo-api/services"
    "github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
    app.Post("/create_student", services.CreateStudent)
    app.Get("/get_student/:studentId", services.GetStudent)
    app.Put("/update_student/:studentId", services.UpdateStudent)
    app.Delete("/delete_student/:studentId", services.DeleteStudent)
    app.Get("/get_all_students", services.GetAllStudent)
}