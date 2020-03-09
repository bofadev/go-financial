package database

import (
	"fmt"

	"github.com/heroku/go-financial/internals/models"
)

type ProviderFake struct {
	reviews []models.Review
	counter uint
}

//	https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
func helpRemove(slice []models.Review, index int) []models.Review {
	return append(slice[:index], slice[index+1:]...)
}

func (p *ProviderFake) GetHelloString() string {
	return "Hello!"
}
func isReviewValid(r *models.Review) bool {

	if r.Type != models.ReviewTypeAsset && r.Type != models.ReviewTypeLiability {
		return false
	}

	return true
}

func (p *ProviderFake) GetLastReviewId() (bool, uint) {

	lastReview := p.reviews[(len(p.reviews) - 1)]
	return true, lastReview.Key
}

func (p *ProviderFake) AddReview(r *models.Review) bool {

	if isReviewValid(r) == false {
		fmt.Printf("Failed to add review. Bad type/kind: %s\n", r.GetString())
		return false
	}

	r.Key = p.counter
	p.counter += 1

	p.reviews = append(p.reviews, *r)

	return true
}
func (p *ProviderFake) DeleteReview(key uint) bool {

	removalIndex := -1
	for i, v := range p.reviews {
		if v.Key == key {
			removalIndex = i
			break
		}
	}

	if removalIndex > -1 {
		fmt.Printf("Deleting review: %s\n", p.reviews[removalIndex].GetString())
		p.reviews = helpRemove(p.reviews, removalIndex)
		return true
	}

	return false
}
func (p *ProviderFake) GetReviews() []models.Review {
	return p.reviews
}

func (p *ProviderFake) GetSumAssets() (bool, float32) {

	var sum float32
	reviews := p.GetReviews()
	for _, v := range reviews {
		if v.Type == models.ReviewTypeAsset {
			sum += v.Balance
		}
	}

	return true, sum
}
func (p *ProviderFake) GetSumLiabilities() (bool, float32) {

	var sum float32
	reviews := p.GetReviews()
	for _, v := range reviews {
		if v.Type == models.ReviewTypeLiability {
			sum += v.Balance
		}
	}

	return true, sum
}

func (p *ProviderFake) Init() bool {

	p.reviews = make([]models.Review, 0)
	p.counter = 1

	var i uint = 1

	//	Assets.
	for ; i < 5; i++ {
		p.AddReview(&models.Review{Type: models.ReviewTypeAsset, Name: fmt.Sprintf("name%d", i), Balance: 200.0})
	}

	//	Liabilities.
	i = 5
	for ; i < 10; i++ {
		p.AddReview(&models.Review{Type: models.ReviewTypeLiability, Name: fmt.Sprintf("name%d", i), Balance: 100.0})
	}

	return true
}
