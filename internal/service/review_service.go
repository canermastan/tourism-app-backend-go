package service

import (
	"errors"
	"github.com/canermastan/teknofest2025-go-backend/internal/model"
	"github.com/canermastan/teknofest2025-go-backend/internal/repository"
)

type ReviewService struct {
	repo *repository.ReviewRepository
}

func NewReviewService(repo *repository.ReviewRepository) *ReviewService {
	return &ReviewService{
		repo: repo,
	}
}

func (s *ReviewService) Create(review *model.Review) error {
	if review.Rating < 1 || review.Rating > 5 {
		return errors.New("rating should be between 1 and 5")
	}

	review.Visibility = false
	return s.repo.Create(review)
}

func (s *ReviewService) GetAll() ([]model.Review, error) {
	return s.repo.GetAll()
}

func (s *ReviewService) GetByPlaceID(placeID int64) ([]model.Review, error) {
	if placeID <= 0 {
		return nil, errors.New("invalid place ID")
	}
	return s.repo.GetByPlaceID(placeID)
}

func (s *ReviewService) GetByPlaceIDAndUserID(placeID, userID int64) ([]model.Review, error) {
	if placeID <= 0 || userID <= 0 {
		return nil, errors.New("invalid place ID or user ID")
	}
	return s.repo.GetByPlaceIDAndUserID(placeID, userID)
}

func (s *ReviewService) GetByID(id int64) (*model.Review, error) {
	if id <= 0 {
		return nil, errors.New("invalid review ID")
	}
	return s.repo.GetByID(id)
}

func (s *ReviewService) Update(review *model.Review) error {
	if review.ID <= 0 {
		return errors.New("invalid review ID")
	}
	if review.Rating < 1 || review.Rating > 5 {
		return errors.New("rating should be between 1 and 5")
	}
	return s.repo.Update(review)
}

func (s *ReviewService) Delete(id int64) error {
	if id <= 0 {
		return errors.New("invalid review ID")
	}
	return s.repo.Delete(id)
}
