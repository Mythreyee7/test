package services

import (
	"context"
	"fiber-mongo-api/config"
	"fiber-mongo-api/models"
	"fiber-mongo-api/responses"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var studentCollection *mongo.Collection = config.GetCollection(config.DB, "students")
var validate = validator.New()

func CreateStudent(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	fmt.Print(ctx)
    var student models.Student
    defer cancel()
  
    if err := c.BodyParser(&student); err != nil {
        return c.Status(http.StatusBadRequest).JSON(responses.StudentResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }
  
    if validationErr := validate.Struct(&student); validationErr != nil {
        return c.Status(http.StatusBadRequest).JSON(responses.StudentResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
    }
  
    newStudent := models.Student{
        Id:          primitive.NewObjectID(),
        Name:        student.Name,
		Student_id:  student.Student_id,
        Register_no: student.Register_no,
        Department:  student.Department,
    }
  
    result, err := studentCollection.InsertOne(ctx, newStudent)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(responses.StudentResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }
  
    return c.Status(http.StatusCreated).JSON(responses.StudentResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

func GetStudent(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    studentId := c.Params("studentId")
    var student models.Student
    defer cancel()
    student_id,_:= strconv.Atoi(studentId)
    fmt.Println(studentId)
    err := studentCollection.FindOne(ctx, bson.M{"student_id":int(student_id)}).Decode(&student)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(responses.StudentResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }
  
    return c.Status(http.StatusOK).JSON(responses.StudentResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": student}})
}

func UpdateStudent(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	studentId := c.Params("studentId")
    var student models.Student
    defer cancel()
    student_id,_:= strconv.Atoi(studentId)
  
    if err := c.BodyParser(&student); err != nil {
        return c.Status(http.StatusBadRequest).JSON(responses.StudentResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }
  
    if validationErr := validate.Struct(&student); validationErr != nil {
        return c.Status(http.StatusBadRequest).JSON(responses.StudentResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
    }
  
    update := bson.M{"name": student.Name, "student_id":student.Student_id, "register_no": student.Register_no, "department": student.Department}
  
    result, err := studentCollection.UpdateOne(ctx, bson.M{"student_id":int(student_id)}, bson.M{"$set": update})
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(responses.StudentResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }
  
    var updatedStudent models.Student
    if result.MatchedCount == 1 {
        err := studentCollection.FindOne(ctx, bson.M{"student_id":int(student_id)}).Decode(&updatedStudent)
        if err != nil {
            return c.Status(http.StatusInternalServerError).JSON(responses.StudentResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
        }
    }
  
    return c.Status(http.StatusOK).JSON(responses.StudentResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updatedStudent}})
}

func DeleteStudent(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    studentId := c.Params("studentId")
    defer cancel()
    student_id,_:= strconv.Atoi(studentId)
  
    result, err := studentCollection.DeleteOne(ctx, bson.M{"student_id":int(student_id)})
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(responses.StudentResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }
  
    if result.DeletedCount < 1 {
        return c.Status(http.StatusNotFound).JSON(
            responses.StudentResponse{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": "User with specified ID not found!"}},
        )
    }
  
    return c.Status(http.StatusOK).JSON(
        responses.StudentResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "User successfully deleted!"}},
    )
}

func GetAllStudent(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	fmt.Println("1",ctx)
    var students []models.Student
    defer cancel()
  
    results, err := studentCollection.Find(ctx, bson.M{})
	fmt.Println(results)
  
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(responses.StudentResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }
  
    defer results.Close(ctx)
    for results.Next(ctx) {
        var singleStudent models.Student
		// res:= results.Decode(&singleStudent)
		// fmt.Println("result", res)
        if err = results.Decode(&singleStudent); 
		err != nil {
            return c.Status(http.StatusInternalServerError).JSON(responses.StudentResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
        }
      
        students = append(students, singleStudent)
    }
  
    return c.Status(http.StatusOK).JSON(
        responses.StudentResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": students}},
    )
}