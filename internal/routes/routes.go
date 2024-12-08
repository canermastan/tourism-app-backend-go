package routes

import (
	"github.com/canermastan/teknofest2025-go-backend/internal/controller"
	"github.com/canermastan/teknofest2025-go-backend/internal/repository"
	"github.com/canermastan/teknofest2025-go-backend/internal/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api")

	// Reviews
	reviewRepository := repository.NewReviewRepository(db)
	reviewService := service.NewReviewService(reviewRepository)
	reviewController := controller.NewReviewController(reviewService)

	review := api.Group("/review")
	review.Post("/save", reviewController.Create)     // POST /api/review/save
	review.Get("/findAll", reviewController.GetAll)   // GET /api/review/find
	review.Get("/find/:id", reviewController.GetByID) // GET /api/review/find/:id
	review.Put("/:id", reviewController.Update)       // PUT /api/review/update/:id
	review.Delete("/:id", reviewController.Delete)    // DELETE /api/review/delete/:id

	// Other routes...

}
