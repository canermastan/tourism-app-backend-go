package repository

import (
	"github.com/canermastan/teknofest2025-go-backend/internal/model"
	"gorm.io/gorm"
)

type ReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) *ReviewRepository {
	return &ReviewRepository{
		db: db,
	}
}

func (r *ReviewRepository) Create(review *model.Review) error {
	return r.db.Create(review).Error
}

func (r *ReviewRepository) GetAll() ([]model.Review, error) {
	var reviews []model.Review
	err := r.db.Find(&reviews).Error
	return reviews, err
}

func (r *ReviewRepository) GetById(id int64) (*model.Review, error) {
	var review model.Review
	err := r.db.First(&review, id).Error
	if err != nil {
		return nil, err
	}
	return &review, nil
}

func (r *ReviewRepository) Update(review *model.Review) error {
	return r.db.Save(review).Error
}

func (r *ReviewRepository) Delete(id int64) error {
	return r.db.Delete(&model.Review{}, id).Error
}
