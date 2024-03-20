package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/RionDsilvaCS/configs"
	"github.com/RionDsilvaCS/models"
	"github.com/RionDsilvaCS/responses"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var jobCollection *mongo.Collection = configs.GetCollection(configs.DB, "jobs")

// var validate = validator.New()

func CreateJob(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var job models.Job
	defer cancel()

	if err := c.BodyParser(&job); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	configs.ValidateJob(c, job)

	newJob := models.Job{
		Id:            primitive.NewObjectID(),
		Job_title:     job.Job_title,
		Job_des:       job.Company,
		Company:       job.Company,
		Location:      job.Location,
		Start_date:    job.Start_date,
		Last_date:     job.Last_date,
		Salary:        job.Salary,
		Recruit_count: job.Recruit_count,
		Type:          job.Type,
		Duration:      job.Duration,
		Applicants:    job.Applicants,
		Skills:        job.Skills,
		User_ID:       job.User_ID,
	}

	result, err := jobCollection.InsertOne(ctx, newJob)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

// func GetAJob(c *fiber.Ctx) error {

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	userId := c.Params("userId")

// 	var user models.User
// 	defer cancel()

// 	print(userId)
// 	objId, _ := primitive.ObjectIDFromHex(userId)
// 	err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
// 	}

// 	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": user}})
// }

func GetAllJobs(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var jobs []models.Job
	defer cancel()

	results, err := jobCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleJob models.Job
		if err = results.Decode(&singleJob); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		jobs = append(jobs, singleJob)
	}

	return c.Status(http.StatusOK).JSON(
		responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": jobs}},
	)
}
