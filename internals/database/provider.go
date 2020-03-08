package database

import (
	"github.com/heroku/go-financial/internals/models"
)

type Provider interface {
	GetHelloString() string
	AddReview(r *models.Review) (bool, uint)
	DeleteReview(key uint) bool
	GetReviews() []models.Review
}
