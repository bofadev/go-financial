package database

import (
	"github.com/heroku/go-financial/internals/models"
)

type Provider interface {
	GetHelloString() string
	AddReview(r *models.Review) (bool, uint)
	DeleteReview(key uint) bool
	SetReviewBalance(key uint, balance float32) bool
	GetReviews() []models.Review
}
