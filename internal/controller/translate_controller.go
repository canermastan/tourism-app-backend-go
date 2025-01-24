package controller

import (
	"bytes"
	"encoding/json"
	"github.com/canermastan/teknofest2025-go-backend/internal/model/dto"
	"github.com/canermastan/teknofest2025-go-backend/internal/response"
	"github.com/gofiber/fiber/v2"
	"io"
	"log"
	"net/http"
)

func TranslateText(ctx *fiber.Ctx) error {
	body := ctx.Body()

	var data map[string]string
	if err := json.Unmarshal(body, &data); err != nil {
		log.Println("JSON deserialize edilemedi:", err)
		return response.ErrorResponse(ctx, fiber.StatusBadRequest, "Gelen cevap JSON formatında değil.")
	}

	text, ok := data["text"]
	if !ok {
		log.Println("Text'e erişilemedi.")
		return response.ErrorResponse(ctx, fiber.StatusBadRequest, "'text' alanı gerekli.")
	}

	url := "https://translate.canermastan.com"
	payload := dto.TranslateRequest{
		Text:         text,
		Source:       "auto",
		Target:       "en",
		Format:       "text",
		Alternatives: 0,
		ApiKey:       "",
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Println("JSON serialize edilemedi:", err)
		return response.ErrorResponse(ctx, fiber.StatusInternalServerError, "Sunucu hatası")
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("API isteği başarısız:", err)
		return response.ErrorResponse(ctx, fiber.StatusInternalServerError, "Çeviri servisine ulaşılamadı")
	}

	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Println("API cevabı okunamadı:", err)
		return response.ErrorResponse(ctx, fiber.StatusInternalServerError, "Çeviri cevabı alınamadı")
	}

	var apiResponse dto.TranslateResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		log.Println("API cevabı JSON parse edilemedi:", err)
		return response.ErrorResponse(ctx, fiber.StatusInternalServerError, "Çeviri cevabı işlenemedi")
	}
	return response.SuccessResponse(ctx, apiResponse, "Çeviri başarılı")
}
