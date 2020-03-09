package database

import (
	"github.com/heroku/go-financial/internals/models"
)

type Provider interface {
	GetHelloString() string
	AddReview(r *models.Review) bool
	DeleteReview(key uint) bool
	GetReviews() []models.Review
	GetSumAssets() (bool, float32)
	GetSumLiabilities() (bool, float32)
	GetLastReviewId() (bool, uint)
	Init() bool
}
