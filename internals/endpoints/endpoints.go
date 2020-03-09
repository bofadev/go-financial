package endpoints

import (
	"encoding/json"
	"fmt"

	"github.com/heroku/go-financial/internals/models"

	"github.com/heroku/go-financial/internals/database"
)

type Endpoints struct {
	db database.Database
}

//	DBG.
func (e *Endpoints) SayHello() {
	fmt.Println("Endpoint Says Mroow!")
}

//	SETUP.
func (e *Endpoints) SetupFake() {

	fmt.Println("Setup with db FAKE")

	fake := &database.ProviderFake{}
	e.db.SetProvider(fake)

	fake.Init()
}
func (e *Endpoints) Setup() {

	fmt.Println("Setup with db REAL")

	psql := &database.ProviderPSQL{}
	e.db.SetProvider(psql)

	if ok := psql.Init(); ok == false {
		panic("Database failed to init")
	}

}

//	ADD/REMOVE.
func (e *Endpoints) AddReview(kind models.ReviewType, name string, balance float32) bool {
	ok := e.db.AddReview(&models.Review{Type: kind, Name: name, Balance: balance})

	if ok == false {

		fmt.Println("Endpoint - could not add review")
		return false
	}

	return true
}
func (e *Endpoints) RemoveReview(key uint) bool {

	if ok := e.db.DeleteReview(key); ok == false {

		fmt.Println("Endpoint - could not delete review")
		return false
	}

	return true
}

//	GET.
func (e *Endpoints) GetReviews() string {

	data, err := json.Marshal(e.db.GetReviews())
	if err != nil {
		fmt.Println("EndPoints - failed to get reviews!")
		return ""
	}
	return string(data)
}

//	Returns the summed total of Assets and Liabilities in the database.
func (e *Endpoints) GetNetWorth() string {

	ok, sumAssets := e.db.GetSumAssets()
	if ok == false {
		fmt.Println("Error getting asset sum from db")
	}
	ok, sumLiabilities := e.db.GetSumLiabilities()
	if ok == false {
		fmt.Println("Error getting liability sum from db")
	}

	return fmt.Sprintf("%.2f", sumAssets-sumLiabilities)
}

//	Returns the summed total of the assets in the database.
func (e *Endpoints) GetAssetsTotal() string {

	ok, sum := e.db.GetSumAssets()
	if ok == false {
		fmt.Println("Error getting asset sum from db")
	}

	return fmt.Sprintf("%.2f", sum)
}

//	Returns the summed total of liabilities in the database.
func (e *Endpoints) GetLiabilitiesTotal() string {

	ok, sum := e.db.GetSumLiabilities()
	if ok == false {
		fmt.Println("Error getting liability sum from db")
	}

	return fmt.Sprintf("%.2f", sum)
}

func (e *Endpoints) GetLastReviewId() (bool, uint) {
	return e.db.GetLastReviewId()
}
