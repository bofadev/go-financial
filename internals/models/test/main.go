package main

import (
	"fmt"

	"github.com/heroku/go-financial/internals/models"
)

func main() {

	r := models.Review{Type: "default", Name: "yorik", Balance: 123.456}
	fmt.Println(r.GetString())
}
