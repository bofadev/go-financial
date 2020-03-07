package main

import (
	"fmt"

	"github.com/heroku/go-financial/internals/models"

	"github.com/heroku/go-financial/internals/endpoints"
)

func main() {

	e := endpoints.Endpoints{}

	e.SetupFake()

	e.SetReviewBalance(1, 5000)
	e.AddReview(models.ReviewTypeAsset, "bobbert", 222.22)
	e.RemoveReview(1)

	fmt.Print(e.GetNetWorth())
	fmt.Print(e.GetAssetsTotal())
	fmt.Print(e.GetLiabilitiesTotal())

	fmt.Printf("json string: %s\n", e.GetReviews())
}
