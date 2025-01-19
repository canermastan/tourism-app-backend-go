package controller

import (
	"github.com/canermastan/teknofest2025-go-backend/internal/model"
	"github.com/canermastan/teknofest2025-go-backend/internal/response"
	"github.com/canermastan/teknofest2025-go-backend/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CollectedChestController struct {
	service *service.CollectedChestService
}

func NewCollectedChestController(service *service.CollectedChestService) *CollectedChestController {
	return &CollectedChestController{
		service: service,
	}
}

func (cc *CollectedChestController) GetByUserID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.ErrorResponse(c, 400, "Geçersiz User ID.")
	}

	collectedChests, err := cc.service.GetByUserID(int64(id))
	if err != nil {
		return response.ErrorResponse(c, 404, "Kayıt bulunamadı.")
	}

	return response.SuccessResponse(c, collectedChests, "Kayıtlar bulundu.")
}

func (cc *CollectedChestController) Create(c *fiber.Ctx) error {
	var collectedChest model.CollectedChest
	if err := c.BodyParser(&collectedChest); err != nil {
		return response.ErrorResponse(c, 400, "Geçersiz veri formatı.")
	}

	validate := validator.New()
	if err := validate.Struct(&collectedChest); err != nil {
		return response.ErrorResponse(c, 400, err.Error())
	}

	if err := cc.service.Create(&collectedChest); err != nil {
		return response.ErrorResponse(c, 500, err.Error())
	}

	return response.SuccessResponse(c, collectedChest, "Kayıt oluşturuldu.")
}

func (cc *CollectedChestController) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.ErrorResponse(c, 400, "Geçersiz ID.")
	}

	var collectedChest model.CollectedChest
	if err := c.BodyParser(&collectedChest); err != nil {
		return response.ErrorResponse(c, 400, "Geçersiz veri formatı.")
	}
	collectedChest.ID = int64(id)

	if err := cc.service.Update(&collectedChest); err != nil {
		return response.ErrorResponse(c, 500, err.Error())
	}

	return response.SuccessResponse(c, collectedChest, "Kayıt güncellendi.")
}

func (cc *CollectedChestController) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.ErrorResponse(c, 400, "Geçersiz ID.")
	}

	if err := cc.service.Delete(int64(id)); err != nil {
		return response.ErrorResponse(c, 500, err.Error())
	}

	return response.SuccessResponse(c, nil, "Kayıt silindi.")
}
