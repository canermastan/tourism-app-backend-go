package repository

import (
	"github.com/canermastan/teknofest2025-go-backend/internal/model"
	"gorm.io/gorm"
)

type CollectedChestRepository struct {
	db *gorm.DB
}

func NewCollectedChestRepository(db *gorm.DB) *CollectedChestRepository {
	return &CollectedChestRepository{
		db: db,
	}
}

func (r *CollectedChestRepository) Create(collectedChest *model.CollectedChest) error {
	return r.db.Create(collectedChest).Error
}

func (r *CollectedChestRepository) Update(collectedChest *model.CollectedChest) error {
	return r.db.Save(collectedChest).Error
}

func (r *CollectedChestRepository) Delete(id int64) error {
	return r.db.Delete(&model.CollectedChest{}, id).Error
}

func (r *CollectedChestRepository) GetByUserID(id int64) ([]model.CollectedChest, error) {
	var collectedChests []model.CollectedChest
	err := r.db.Where("user_id = ?", id).Find(&collectedChests).Error
	if err != nil {
		return nil, err
	}
	return collectedChests, nil
}
