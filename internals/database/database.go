package database

import "github.com/heroku/go-financial/internals/models"

type Database struct {
	provider Provider
}

func (d *Database) SetProvider(p Provider) {
	d.provider = p
}

func (d *Database) GetHelloString() string {
	if d.provider == nil {
		panic("No provider set in database!")
	}
	return d.provider.GetHelloString()
}

func (d *Database) AddReview(r *models.Review) bool {
	if d.provider == nil {
		panic("No provider set in database!")
	}
	return d.provider.AddReview(r)
}
func (d *Database) DeleteReview(key uint) bool {
	if d.provider == nil {
		panic("No provider set in database!")
	}
	return d.provider.DeleteReview(key)
}

func (d *Database) GetReviews() []models.Review {
	if d.provider == nil {
		panic("No provider set in database!")
	}
	return d.provider.GetReviews()
}
func (d *Database) GetSumAssets() (bool, float32) {
	return d.provider.GetSumAssets()
}
func (d *Database) GetSumLiabilities() (bool, float32) {
	return d.provider.GetSumLiabilities()
}

func (d *Database) GetLastReviewId() (bool, uint) {
	return d.provider.GetLastReviewId()
}
