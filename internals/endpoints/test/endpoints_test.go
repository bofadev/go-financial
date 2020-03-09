package main

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/heroku/go-financial/internals/models"

	"github.com/heroku/go-financial/internals/endpoints"
)

func TestSetupFake(t *testing.T) {

	e := endpoints.Endpoints{}
	e.SetupFake()

	reviews := e.GetReviews()
	if reviews == "" {
		t.Error("Reviews string empty")
	}
}

func TestAddReview(t *testing.T) {

	e := endpoints.Endpoints{}
	e.SetupFake()

	now := time.Now()
	rName := fmt.Sprintf("Name%d", now.Unix())

	foundIndexBefore := strings.Index(e.GetReviews(), rName)
	if foundIndexBefore != -1 {
		t.Error("Test poorly designed. Entry was already in database.")
	}

	e.AddReview(models.ReviewTypeAsset, rName, 3456.0)
	foundIndexAfter := strings.Index(e.GetReviews(), rName)
	if foundIndexAfter == -1 {
		t.Error("Could not find added review")
	}
	//fmt.Printf("Found added review [string index:%d] [review:%s]\n", foundIndexAfter, e.GetReviews())
}

func TestRemoveReview(t *testing.T) {

	e := endpoints.Endpoints{}
	e.SetupFake()

	now := time.Now()
	rName := fmt.Sprintf("Name%d", now.Unix())

	foundIndexBefore := strings.Index(e.GetReviews(), rName)
	if foundIndexBefore != -1 {
		t.Error("Test poorly designed. Entry was already in database.")
	}

	ok := e.AddReview(models.ReviewTypeAsset, rName, 3456.0)
	if ok == false {
		t.Error("AddReview() returned an error")
	}
	foundIndexAfterAdding := strings.Index(e.GetReviews(), rName)
	if foundIndexAfterAdding == -1 {
		t.Error("Could not find added review")
	}
	//fmt.Printf("Found added review [string index:%d] [review:%s]\n", foundIndexAfter, e.GetReviews())

	ok, key := e.GetLastReviewId()
	if ok == false {
		t.Error("Could not get last review id")
	}
	ok = e.RemoveReview(key)
	if ok == false {
		t.Error("RemoveReview() returned an error")
	}
	foundIndexAfterRemoving := strings.Index(e.GetReviews(), rName)
	if foundIndexAfterRemoving != -1 {
		t.Error("Found removed review")
	}
}

func TestGetReviews(t *testing.T) {

	e := endpoints.Endpoints{}
	e.SetupFake()

	reviews := e.GetReviews()
	if reviews == "" {
		t.Error("Got an empty string. Expected at least the default reviews.")
	}
}

func TestGetNetWorth(t *testing.T) {

	e := endpoints.Endpoints{}
	e.SetupFake()

	netWorth := e.GetNetWorth()
	if netWorth == "" {
		t.Error("Got an empty string.")
	}
}

func TestGetAssetsTotal(t *testing.T) {

	e := endpoints.Endpoints{}
	e.SetupFake()

	assetsTotal := e.GetAssetsTotal()
	if assetsTotal == "" {
		t.Error("Got an empty string.")
	}
}
func TestGetLiabilitiesTotal(t *testing.T) {

	e := endpoints.Endpoints{}
	e.SetupFake()

	liabilitiesTotal := e.GetLiabilitiesTotal()
	if liabilitiesTotal == "" {
		t.Error("Got an empty string.")
	}
}
