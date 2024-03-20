package routes

import (
	"github.com/RionDsilvaCS/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Post("/user", controllers.CreateUser)
	app.Get("/user/:userId", controllers.GetAUser)
	app.Get("/users", controllers.GetAllUsers)
}

// func TestimonialRoute(app *fiber.App) {
// 	app.Post("/testimonial", controllers.CreateTestimonial)
// 	app.Get("/testimonial/:testimonialId", controllers.GetATestimonial)
// 	app.Get("/testimonials", controllers.GetAllTestimonials)
// }

func JobRoute(app *fiber.App) {
	app.Post("/job", controllers.CreateJob)
	// app.Get("/job/:jobId", controllers.GetAJob)
	app.Get("/jobs", controllers.GetAllJobs)
}

// func ApplicationRoute(app *fiber.App) {
// 	app.Post("/application", controllers.CreateApplication)
// 	app.Get("/application/:applicationId", controllers.GetAApplication)
// 	app.Get("/applications", controllers.GetAllApplications)
// }

// app.Put("/testimonial/:testimonialId", controllers.EditAUser)
// app.Delete("/testimonial/:testimonialId", controllers.DeleteAUser)
