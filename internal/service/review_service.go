package service

import (
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
	if review.Rating > 5 {
		// TODO: return &model.ValidationError{Message: "Rating should be between 1 and 5."}
	}

	review.Visibility = false
	return s.repo.Create(review)
}

func (s *ReviewService) GetAll() ([]model.Review, error) {
	return s.repo.GetAll()
}

func (s *ReviewService) GetByID(id int64) (*model.Review, error) {
	return s.repo.GetById(id)
}

func (s *ReviewService) Update(review *model.Review) error {
	return s.repo.Update(review)
}

func (s *ReviewService) Delete(id int64) error {
	return s.repo.Delete(id)
}
