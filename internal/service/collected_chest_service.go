package service

import (
	"github.com/canermastan/teknofest2025-go-backend/internal/model"
	"github.com/canermastan/teknofest2025-go-backend/internal/repository"
)

type CollectedChestService struct {
	repository *repository.CollectedChestRepository
}

func NewCollectedChestService(repository *repository.CollectedChestRepository) *CollectedChestService {
	return &CollectedChestService{
		repository: repository,
	}
}

func (s *CollectedChestService) Create(collectedChest *model.CollectedChest) error {
	return s.repository.Create(collectedChest)
}

func (s *CollectedChestService) Update(collectedChest *model.CollectedChest) error {
	return s.repository.Update(collectedChest)
}

func (s *CollectedChestService) Delete(id int64) error {
	return s.repository.Delete(id)
}

func (s *CollectedChestService) GetByUserID(id int64) ([]model.CollectedChest, error) {
	return s.repository.GetByUserID(id)
}
