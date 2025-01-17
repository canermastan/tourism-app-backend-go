package service

import (
	"github.com/canermastan/teknofest2025-go-backend/internal/model"
	"github.com/canermastan/teknofest2025-go-backend/internal/repository"
)

type ChestService struct {
	repo *repository.ChestRepository
}

func NewChestService(chestRepository *repository.ChestRepository) *ChestService {
	return &ChestService{
		repo: chestRepository,
	}
}

func (c *ChestService) Create(chest *model.Chest) error {
	if chest.PlaceID <= 0 {
		return &model.ValidationError{Message: "Place ID must be a positive integer and cannot be zero or undefined."}
	}

	return c.repo.Create(chest)
}

func (c *ChestService) Update(chest *model.Chest) error {
	return c.repo.Update(chest)
}

func (c *ChestService) Delete(id int64) error {
	return c.repo.Delete(id)
}

func (c *ChestService) GetByID(id int64) (*model.Chest, error) {
	return c.repo.GetByID(id)
}

func (c *ChestService) GetAll() ([]model.Chest, error) {
	return c.repo.GetAll()
}
