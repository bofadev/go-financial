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

func (p *ProviderFake) AddReview(r *models.Review) (bool, uint) {

	if isReviewValid(r) == false {
		fmt.Printf("Failed to add review. Bad type/kind: %s\n", r.GetString())
		return false, 0
	}

	r.Key = p.counter
	p.counter += 1

	p.reviews = append(p.reviews, *r)

	return true, r.Key
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
func (p *ProviderFake) SetReviewBalance(key uint, balance float32) bool {

	for i, v := range p.reviews {
		if v.Key == key {
			p.reviews[i].Balance = balance
			fmt.Printf("SetReviewBalance - Success: %s\n", p.reviews[i].GetString())
			return true
		}
	}

	fmt.Printf("SetReviewBalance - could not find review w/ key %d", key)
	return false
}
func (p *ProviderFake) GetReviews() []models.Review {
	return p.reviews
}

func (p *ProviderFake) Populate() {

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
}
