package controller

import (
	"github.com/canermastan/teknofest2025-go-backend/internal/model"
	"github.com/canermastan/teknofest2025-go-backend/internal/response"
	"github.com/canermastan/teknofest2025-go-backend/internal/service"
	"github.com/gofiber/fiber/v2"
)

type ChestController struct {
	service *service.ChestService
}

func NewChestController(service *service.ChestService) *ChestController {
	return &ChestController{
		service: service,
	}
}

func (c *ChestController) Create(ctx *fiber.Ctx) error {
	var chest model.Chest
	if err := ctx.BodyParser(&chest); err != nil {
		return response.ErrorResponse(ctx, 400, err.Error())
	}
	err := c.service.Create(&chest)
	if err != nil {
		return response.ErrorResponse(ctx, 500, err.Error())
	}
	return response.SuccessResponse(ctx, 200, "Kayıt oluşturuldu.")
}

func (c *ChestController) Update(ctx *fiber.Ctx) error {
	var chest model.Chest
	if err := ctx.BodyParser(&chest); err != nil {
		return response.ErrorResponse(ctx, 400, err.Error())
	}

	id, err := ctx.ParamsInt("id")
	if err != nil {
		return response.ErrorResponse(ctx, 400, err.Error())
	}

	chest.ID = int64(id)
	err = c.service.Update(&chest)
	if err != nil {
		return response.ErrorResponse(ctx, 500, err.Error())
	}
	return response.SuccessResponse(ctx, 200, "Kayıt güncellendi.")
}

func (c *ChestController) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return response.ErrorResponse(ctx, 400, err.Error())
	}
	if err := c.service.Delete(int64(id)); err != nil {
		return response.ErrorResponse(ctx, 500, err.Error())
	}

	return response.SuccessResponse(ctx, 200, "Kayıt silindi.")
}

func (c *ChestController) GetById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return response.ErrorResponse(ctx, 400, err.Error())
	}

	chest, err := c.service.GetByID(int64(id))
	if err != nil {
		return response.ErrorResponse(ctx, 404, "Kayıt bulunamadı.")
	}

	return response.SuccessResponse(ctx, chest, "Kayıt getirildi.")
}

func (c *ChestController) GetAll(ctx *fiber.Ctx) error {
	chests, err := c.service.GetAll()
	if err != nil {
		return response.ErrorResponse(ctx, 500, "Kayıtlar getirilemedi.")
	}

	return response.SuccessResponse(ctx, chests, "Kayıtlar getirildi.")
}
