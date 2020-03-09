package main

import (
	"github.com/heroku/go-financial/internals/database"
)

func main() {

	p := &database.ProviderPSQL{}

	p.Init()
}
