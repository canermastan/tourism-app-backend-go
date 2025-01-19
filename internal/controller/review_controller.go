package controller

import (
	"github.com/canermastan/teknofest2025-go-backend/internal/model"
	"github.com/canermastan/teknofest2025-go-backend/internal/response"
	"github.com/canermastan/teknofest2025-go-backend/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ReviewController struct {
	service *service.ReviewService
}

func NewReviewController(service *service.ReviewService) *ReviewController {
	return &ReviewController{
		service: service,
	}
}

func (rc *ReviewController) GetAll(c *fiber.Ctx) error {
	reviews, err := rc.service.GetAll()
	if err != nil {
		return response.ErrorResponse(c, 400, "Kayıt bulunamadı.")
	}
	return response.SuccessResponse(c, reviews, "Kayıtlar bulundu.")
}

func (rc *ReviewController) GetByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.ErrorResponse(c, 400, "Geçersiz ID.")
	}

	review, err := rc.service.GetByID(int64(id))
	if err != nil {
		return response.ErrorResponse(c, 404, "Kayıt bulunamadı.")
	}

	return response.SuccessResponse(c, review, "Kayıt bulundu.")
}

func (rc *ReviewController) GetByPlaceID(c *fiber.Ctx) error {
	placeID, err := c.ParamsInt("place_id")
	if err != nil {
		return response.ErrorResponse(c, 400, "Geçersiz Place ID.")
	}

	reviews, err := rc.service.GetByPlaceID(int64(placeID))
	if err != nil {
		return response.ErrorResponse(c, 404, "Kayıt bulunamadı.")
	}

	return response.SuccessResponse(c, reviews, "Kayıtlar bulundu.")
}

func (rc *ReviewController) GetByPlaceIDAndUserID(c *fiber.Ctx) error {
	placeID, err := c.ParamsInt("place_id")
	if err != nil {
		return response.ErrorResponse(c, 400, "Geçersiz Place ID.")
	}

	userID, err := c.ParamsInt("user_id")
	if err != nil {
		return response.ErrorResponse(c, 400, "Geçersiz User ID.")
	}

	reviews, err := rc.service.GetByPlaceIDAndUserID(int64(placeID), int64(userID))
	if err != nil {
		return response.ErrorResponse(c, 404, "Kayıt bulunamadı.")
	}

	return response.SuccessResponse(c, reviews, "Kayıtlar bulundu.")
}

func (rc *ReviewController) Create(c *fiber.Ctx) error {
	var review model.Review
	if err := c.BodyParser(&review); err != nil {
		return response.ErrorResponse(c, 400, "Geçersiz veri formatı.")
	}

	validate := validator.New()
	if err := validate.Struct(&review); err != nil {
		return response.ErrorResponse(c, 400, err.Error())
	}

	if err := rc.service.Create(&review); err != nil {
		return response.ErrorResponse(c, 500, err.Error())
	}

	return response.SuccessResponse(c, review, "Kayıt oluşturuldu.")
}

func (rc *ReviewController) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.ErrorResponse(c, 400, "Geçersiz ID.")
	}

	var review model.Review
	if err := c.BodyParser(&review); err != nil {
		return response.ErrorResponse(c, 400, "Geçersiz veri formatı.")
	}
	review.ID = int64(id)

	if err := rc.service.Update(&review); err != nil {
		return response.ErrorResponse(c, 500, err.Error())
	}

	return response.SuccessResponse(c, review, "Kayıt güncellendi.")
}

func (rc *ReviewController) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.ErrorResponse(c, 400, "Geçersiz ID.")
	}

	if err := rc.service.Delete(int64(id)); err != nil {
		return response.ErrorResponse(c, 500, err.Error())
	}

	return response.SuccessResponse(c, nil, "Kayıt silindi.")
}
