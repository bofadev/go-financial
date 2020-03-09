package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/heroku/go-financial/internals/models"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/heroku/go-financial/internals/endpoints"
)

type Response struct {
	Success          bool
	Key              string
	NetWorth         string
	TotalAssets      string
	TotalLiabilities string
	Message          string
}

func getResponseJsonString(success bool, key string, netWorth string, totalAssets string, totalLiabilities string, message string) string {
	r := Response{Success: success, Key: key, NetWorth: netWorth, TotalAssets: totalAssets, TotalLiabilities: totalLiabilities, Message: message}

	b, err := json.Marshal(r)
	if err != nil {
		fmt.Println("Error - could not marshal response into json")
		return ""
	}

	rString := string(b)
	fmt.Println(rString)

	return rString
}

func main() {

	r := gin.Default()

	//	Serve our content ------------------------
	r.Use(static.Serve("/", static.LocalFile("./web/build", true)))

	//	Connect endpoints -----------------------------
	e := endpoints.Endpoints{}
	//e.SetupFake()
	e.Setup()
	const responseCodeOK int = 200
	const responseCodeFailure int = 500

	r.GET("/reviews/get", func(c *gin.Context) {
		c.String(responseCodeOK, e.GetReviews())
	})
	r.GET("/review/add", func(c *gin.Context) {

		params := c.Request.URL.Query()
		if params == nil || params["type"] == nil || params["name"] == nil || params["balance"] == nil {
			rString := getResponseJsonString(false, "", "", "", "", "Failure - tried to add review without supplying all parameters necessary")
			c.String(responseCodeFailure, rString)
		}

		var kind models.ReviewType = models.ReviewType(params.Get("type"))
		name := params.Get("name")
		balance, err := strconv.ParseFloat(params.Get("balance"), 32)

		if err != nil {
			rString := getResponseJsonString(false, "", "", "", "", "Failed to parseFloat balance")
			c.String(responseCodeFailure, rString)
			return
		}

		ok := e.AddReview(kind, name, float32(balance))

		if ok == false {
			rString := getResponseJsonString(false, "", "", "", "", "Failed to add review")
			c.String(responseCodeFailure, rString)
			return
		}

		rString := getResponseJsonString(true, "", "", "", "", "Added review")
		c.String(responseCodeOK, rString)
	})
	r.GET("/review/delete", func(c *gin.Context) {

		params := c.Request.URL.Query()
		if params == nil || params["key"] == nil {
			rString := getResponseJsonString(false, "", "", "", "", "Failure - tried to delete review without supplying all parameters necessary")
			c.String(responseCodeFailure, rString)
		}

		key64, err := strconv.ParseUint(params.Get("key"), 10, 32)
		if err != nil {
			rString := getResponseJsonString(false, "", "", "", "", "Failed to parseUint key")
			c.String(responseCodeFailure, rString)
			return
		}

		key32 := uint(key64)
		if e.RemoveReview(key32) == false {
			rString := getResponseJsonString(false, "", "", "", "", "Failed to remove review")
			c.String(responseCodeFailure, rString)
			return
		}

		rString := getResponseJsonString(true, string(key32), "", "", "", "Success - removed review")
		c.String(responseCodeOK, rString)
	})
	r.GET("/getTotals", func(c *gin.Context) {
		netWorth := e.GetNetWorth()
		assetsTotal := e.GetAssetsTotal()
		liabilitiesTotal := e.GetLiabilitiesTotal()

		rString := getResponseJsonString(true, "", netWorth, assetsTotal, liabilitiesTotal, "Success")
		c.String(responseCodeOK, rString)
	})

	// Listen and Serve -------------------
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	r.Run(":" + port)
}
