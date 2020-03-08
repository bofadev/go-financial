package main

import (
	"testing"

	"github.com/heroku/go-financial/internals/database"
	"github.com/heroku/go-financial/internals/models"
)

func TestSetProvider(t *testing.T) {

	//	Setup.
	db := database.Database{}
	provider := &database.ProviderFake{}
	db.SetProvider(provider)

	//	Check for a message.
	if db.GetHelloString() == "" {
		t.Error("Could not get message from provider")
	}
}

func TestGetHelloString(t *testing.T) {

	//	Setup.
	db := database.Database{}
	provider := &database.ProviderFake{}
	db.SetProvider(provider)

	//	Check for a message.
	if db.GetHelloString() == "" {
		t.Error("Could not get message from provider")
	}
}

func TestAddReview(t *testing.T) {

	//	Setup.
	db := database.Database{}
	provider := &database.ProviderFake{}
	db.SetProvider(provider)

	//	Add and confirm it added.
	r := &models.Review{Type: models.ReviewTypeAsset, Name: "MyReview", Balance: 19282.0}
	ok, _ := db.AddReview(r)
	if ok == false {
		t.Error("AddReview() returned ok FALSE")
	}
}

func TestDeleteReview(t *testing.T) {

	//	Setup.
	db := database.Database{}
	provider := &database.ProviderFake{}
	db.SetProvider(provider)

	//	Add a review.
	r := &models.Review{Type: models.ReviewTypeAsset, Name: "MyReview", Balance: 19282.0}
	ok, keyAdded := db.AddReview(r)
	if ok == false {
		t.Error("AddReview() returned ok FALSE")
	}

	//	Confirm it added.
	found := false
	for _, v := range db.GetReviews() {
		if v.Key == keyAdded {
			found = true
			break
		}
	}
	if found == false {
		t.Error("Could not find added review")
	}

	//	Remove and confirm it removed.
	ok = db.DeleteReview(keyAdded)
	if ok == false {
		t.Error("DeleteReview() returned ok FALSE")
	}
	for _, v := range db.GetReviews() {
		if v.Key == keyAdded {
			if found == true {
				t.Error("Still found review after supposed removal")
			}
		}
	}
}

func TestGetReviews(t *testing.T) {

	//	Setup.
	provider := &database.ProviderFake{}
	db := database.Database{}
	db.SetProvider(provider)
	provider.Populate()

	//	Check that we have reviews.
	if db.GetReviews() == nil {
		t.Error("Did not find any reviews")
	}
}