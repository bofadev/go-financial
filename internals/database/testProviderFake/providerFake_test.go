package main

import (
	"testing"

	"github.com/heroku/go-financial/internals/database"
	"github.com/heroku/go-financial/internals/models"
)

func TestGetHelloString(t *testing.T) {

	//	Setup.
	p := database.ProviderFake{}

	//	Check for a message.
	if p.GetHelloString() == "" {
		t.Error("Got an empty string")
	}
}

func TestAddReview(t *testing.T) {

	//	Setup.
	p := database.ProviderFake{}

	//	Add and confirm it added.
	r := &models.Review{Type: models.ReviewTypeAsset, Name: "MyReview", Balance: 19282.0}
	ok, _ := p.AddReview(r)
	if ok == false {
		t.Error("AddReview() returned ok FALSE")
	}
}

func TestDeleteReview(t *testing.T) {

	//	Setup.
	p := database.ProviderFake{}

	//	Add a review.
	r := &models.Review{Type: models.ReviewTypeAsset, Name: "MyReview", Balance: 19282.0}
	ok, keyAdded := p.AddReview(r)
	if ok == false {
		t.Error("AddReview() returned ok FALSE")
	}

	//	Confirm it added.
	found := false
	for _, v := range p.GetReviews() {
		if v.Key == keyAdded {
			found = true
			break
		}
	}
	if found == false {
		t.Error("Could not find added review")
	}

	//	Remove and confirm it removed.
	ok = p.DeleteReview(keyAdded)
	if ok == false {
		t.Error("DeleteReview() returned ok FALSE")
	}
	for _, v := range p.GetReviews() {
		if v.Key == keyAdded {
			if found == true {
				t.Error("Still found review after supposed removal")
			}
		}
	}
}

func TestGetReviews(t *testing.T) {

	//	Setup.
	p := database.ProviderFake{}

	//	Add and confirm it added.
	r := &models.Review{Type: models.ReviewTypeAsset, Name: "MyReview", Balance: 19282.0}
	ok, _ := p.AddReview(r)
	if ok == false {
		t.Error("AddReview() returned ok FALSE")
	}

	//	Check that we have reviews.
	if p.GetReviews() == nil {
		t.Error("Did not find any reviews")
	}
}

func TestPopulate(t *testing.T) {

	//	Setup.
	p := database.ProviderFake{}
	p.Populate()

	//	Check that we have reviews.
	if p.GetReviews() == nil {
		t.Error("Did not find any reviews")
	}

}
