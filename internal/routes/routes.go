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
	review.Post("/save", reviewController.Create)                                                // POST /api/review/save
	review.Get("/findAll", reviewController.GetAll)                                              // GET /api/review/findAll
	review.Get("/find/:id", reviewController.GetByID)                                            // GET /api/review/find/:id
	review.Get("/findByPlace/:place_id", reviewController.GetByPlaceID)                          // GET /api/review/findByPlace/:place_id
	review.Get("/findByPlaceAndUser/:place_id/:user_id", reviewController.GetByPlaceIDAndUserID) // GET /api/review/findByPlaceAndUser/:place_id/:user_id
	review.Put("/update/:id", reviewController.Update)                                           // PUT /api/review/update/:id
	review.Delete("/delete/:id", reviewController.Delete)                                        // DELETE /api/review/delete/:id

	// Other routes...
	chestRepository := repository.NewChestRepository(db)
	chestService := service.NewChestService(chestRepository)
	chestController := controller.NewChestController(chestService)

	chest := api.Group("/chest")
	chest.Post("/create", chestController.Create)       // POST /api/chest/create
	chest.Put("/update/:id", chestController.Update)    // PUT /api/chest/update/:id
	chest.Delete("/delete/:id", chestController.Delete) // DELETE /api/chest/delete/:id
	chest.Get("/find/:id", chestController.GetById)     // GET /api/chest/find/:id
	chest.Get("/findAll", chestController.GetAll)       // GET /api/chest/findAll
}
