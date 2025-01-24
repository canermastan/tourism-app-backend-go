package test

import (
	"github.com/canermastan/teknofest2025-go-backend/internal/repository"
	"github.com/canermastan/teknofest2025-go-backend/internal/utils"
	"log"
	"os"
	"testing"
	"time"

	"github.com/canermastan/teknofest2025-go-backend/internal/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	testDB        *gorm.DB
	reviewRepo    *repository.ReviewRepository
	sampleReview1 model.Review
	sampleReview2 model.Review
	sampleReview3 model.Review
)

func TestMain(m *testing.M) {
	var err error

	testDB, err = utils.SetupTestDB()
	if err != nil {
		log.Fatalf("Failed to set up test database: %v", err)
	}

	err = testDB.AutoMigrate(&model.Review{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	reviewRepo = repository.NewReviewRepository(testDB)

	sampleReview1 = model.Review{
		UserID:     1,
		PlaceID:    101,
		Rating:     5,
		Comment:    "Excellent place!",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Visibility: true,
	}

	sampleReview2 = model.Review{
		UserID:     2,
		PlaceID:    101,
		Rating:     4,
		Comment:    "Very good, but can improve.",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Visibility: true,
	}

	sampleReview3 = model.Review{
		UserID:     1,
		PlaceID:    102,
		Rating:     3,
		Comment:    "Average experience.",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Visibility: false,
	}

	code := m.Run()

	err = utils.CleanupTestDB(testDB)
	if err != nil {
		log.Printf("Failed to clean up test database: %v", err)
	}

	os.Exit(code)
}

func setupTestData(t *testing.T) {
	err := testDB.Exec("TRUNCATE TABLE reviews RESTART IDENTITY CASCADE;").Error
	assert.NoError(t, err)

	err = testDB.Create(&sampleReview1).Error
	assert.NoError(t, err)
	err = testDB.Create(&sampleReview2).Error
	assert.NoError(t, err)
	err = testDB.Create(&sampleReview3).Error
	assert.NoError(t, err)
}

func TestCreate(t *testing.T) {
	setupTestData(t)
	newReview := model.Review{
		UserID:     3,
		PlaceID:    103,
		Rating:     4,
		Comment:    "Good place!",
		Visibility: true,
	}

	err := reviewRepo.Create(&newReview)

	assert.NoError(t, err)
	assert.NotZero(t, newReview.ID)

	var fetched model.Review
	err = testDB.First(&fetched, newReview.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, newReview.UserID, fetched.UserID)
	assert.Equal(t, newReview.PlaceID, fetched.PlaceID)
	assert.Equal(t, newReview.Rating, fetched.Rating)
	assert.Equal(t, newReview.Comment, fetched.Comment)
	assert.Equal(t, newReview.Visibility, fetched.Visibility)
}

func TestGetAll(t *testing.T) {
	setupTestData(t)

	reviews, err := reviewRepo.GetAll()

	assert.NoError(t, err)
	assert.Len(t, reviews, 3)
}

func TestGetByPlaceID(t *testing.T) {
	setupTestData(t)
	placeID := int64(101)

	reviews, err := reviewRepo.GetByPlaceID(placeID)

	assert.NoError(t, err)
	assert.Len(t, reviews, 2)
	for _, review := range reviews {
		assert.Equal(t, placeID, int64(review.PlaceID))
	}
}

func TestGetByPlaceIDAndUserID(t *testing.T) {
	setupTestData(t)
	placeID := int64(101)
	userID := int64(1)

	reviews, err := reviewRepo.GetByPlaceIDAndUserID(placeID, userID)

	assert.NoError(t, err)
	assert.Len(t, reviews, 1)
	assert.Equal(t, placeID, int64(reviews[0].PlaceID))
	assert.Equal(t, userID, reviews[0].UserID)
}

func TestGetByID(t *testing.T) {
	setupTestData(t)
	existingID := sampleReview1.ID
	nonExistingID := int64(999)

	review, err := reviewRepo.GetByID(existingID)
	assert.NoError(t, err)
	assert.NotNil(t, review)
	assert.Equal(t, sampleReview1.ID, review.ID)

	review, err = reviewRepo.GetByID(nonExistingID)
	assert.NoError(t, err)
	assert.Nil(t, review)
}

func TestUpdate(t *testing.T) {
	setupTestData(t)
	review, err := reviewRepo.GetByID(sampleReview1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, review)

	review.Rating = 4
	review.Comment = "Updated comment."

	err = reviewRepo.Update(review)

	assert.NoError(t, err)

	updatedReview, err := reviewRepo.GetByID(review.ID)
	assert.NoError(t, err)
	assert.NotNil(t, updatedReview)
	assert.Equal(t, byte(4), updatedReview.Rating)
	assert.Equal(t, "Updated comment.", updatedReview.Comment)
}

func TestDelete(t *testing.T) {
	setupTestData(t)
	reviewID := sampleReview2.ID

	err := reviewRepo.Delete(reviewID)

	assert.NoError(t, err)

	deletedReview, err := reviewRepo.GetByID(reviewID)
	assert.NoError(t, err)
	assert.Nil(t, deletedReview)
}

func TestSoftDelete(t *testing.T) {
	setupTestData(t)
	reviewID := sampleReview3.ID

	err := reviewRepo.Delete(reviewID)

	assert.NoError(t, err)

	var deleted model.Review
	err = testDB.Unscoped().First(&deleted, reviewID).Error
	assert.NoError(t, err)
	assert.NotNil(t, deleted.DeletedAt)
}
