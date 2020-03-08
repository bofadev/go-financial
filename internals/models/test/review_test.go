package main

import (
	"testing"

	"github.com/heroku/go-financial/internals/models"
)

func TestGetString(t *testing.T) {

	r := models.Review{Type: "Asset", Name: "FakeName", Balance: 12345}
	msg := r.GetString()

	if msg == "" {
		t.Error("Received empty string")
	}
}

func TestCopyFrom(t *testing.T) {

	r1 := models.Review{Type: models.ReviewTypeAsset, Name: "FakeName", Balance: 12345}
	r2 := models.Review{}

	r2.CopyFrom(&r1)

	if r1.Type != r2.Type {
		t.Error("Copied review Type's not the same")
	}
	if r1.Name != r2.Name {
		t.Error("Copied review Name's not the same")
	}
	if r1.Balance != r2.Balance {
		t.Error("Copied review Balance's not the same")
	}
}

func TestSetTypeAsset(t *testing.T) {

	r := models.Review{Type: models.ReviewTypeLiability}

	r.SetTypeAsset()

	if r.Type != models.ReviewTypeAsset {
		t.Error("Review type not Asset")
	}
}
func TestSetTypeLiability(t *testing.T) {

	r := models.Review{Type: models.ReviewTypeAsset}

	r.SetTypeLiability()

	if r.Type != models.ReviewTypeLiability {
		t.Error("Review type not Liability")
	}
}

func TestIsAsset(t *testing.T) {

	r := models.Review{Type: models.ReviewTypeAsset}

	if r.IsAsset() == true && r.Type != models.ReviewTypeAsset {
		t.Error("Review type is not asset but IsAsset() reports that it is")
	}
}

func TestIsLiability(t *testing.T) {

	r := models.Review{Type: models.ReviewTypeLiability}

	if r.IsLiability() == true && r.Type != models.ReviewTypeLiability {
		t.Error("Review type is not liability but IsLiability() reports that it is")
	}
}
