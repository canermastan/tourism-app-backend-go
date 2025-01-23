package repository

import (
	"errors"
	"github.com/canermastan/teknofest2025-go-backend/internal/model"
	"gorm.io/gorm"
)

type ChestRepository struct {
	db *gorm.DB
}

func NewChestRepository(db *gorm.DB) *ChestRepository {
	return &ChestRepository{
		db: db,
	}
}

func (r *ChestRepository) Create(chest *model.Chest) error { return r.db.Create(chest).Error }

func (r *ChestRepository) Update(chest *model.Chest) error {
	return r.db.Save(chest).Error
}

func (r *ChestRepository) Delete(id int64) error {
	return r.db.Delete(&model.Chest{}, id).Error
}

func (r *ChestRepository) GetByID(id int64) (*model.Chest, error) {
	var chest model.Chest
	err := r.db.First(&chest, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &chest, nil
}

func (r *ChestRepository) GetAll() ([]model.Chest, error) {
	var chests []model.Chest
	err := r.db.Find(&chests).Error
	return chests, err
}
