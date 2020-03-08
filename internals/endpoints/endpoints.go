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
	fake := &database.ProviderFake{}
	e.db.SetProvider(fake)

	fake.Populate()
}

//	ADD/REMOVE.
func (e *Endpoints) AddReview(kind models.ReviewType, name string, balance float32) (bool, uint) {
	ok, val := e.db.AddReview(&models.Review{Type: kind, Name: name, Balance: balance})

	if ok == false {

		fmt.Println("Endpoint - could not add review")
		return false, 0
	}

	return true, val
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

	var val float32 = 0
	for _, v := range e.db.GetReviews() {
		if v.IsAsset() {
			val += v.Balance
		} else if v.IsLiability() {
			val -= v.Balance
		}
	}
	return fmt.Sprintf("%.2f", val)
}

//	Returns the summed total of the assets in the database.
func (e *Endpoints) GetAssetsTotal() string {
	var val float32 = 0
	for _, v := range e.db.GetReviews() {
		if v.IsAsset() {
			val += v.Balance
		}
	}
	return fmt.Sprintf("%.2f", val)
}

//	Returns the summed total of liabilities in the database.
func (e *Endpoints) GetLiabilitiesTotal() string {
	var val float32 = 0
	for _, v := range e.db.GetReviews() {
		if v.IsLiability() {
			val += v.Balance
		}
	}
	return fmt.Sprintf("%.2f", val)
}
