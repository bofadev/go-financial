package main

import (
	"fmt"

	"github.com/heroku/go-financial/internals/database"
)

func main() {

	fakeProvider := &database.ProviderFake{}
	db := database.Database{}
	db.SetProvider(fakeProvider)

	fakeProvider.Populate()

	for _, r := range db.GetReviews() {
		fmt.Println(r.GetString())
	}
}
