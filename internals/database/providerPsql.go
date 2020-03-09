package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/heroku/go-financial/internals/models"
	_ "github.com/lib/pq"
)

type ProviderPSQL struct {
	db *sql.DB
}

func (p *ProviderPSQL) GetHelloString() string {
	return "Hello!"
}
func (p *ProviderPSQL) AddReview(r *models.Review) bool {

	//	Add an item.
	cmdInsert := fmt.Sprintf("INSERT INTO reviews (type, name, balance) VALUES ('%s','%s',%f)",
		r.Type,
		r.Name,
		r.Balance)
	if _, err := p.db.Exec(cmdInsert); err != nil {
		fmt.Printf("Error adding item to db: %q\n", err)
		return false
	}
	// fmt.Println("Added item to database")

	return true
}
func (p *ProviderPSQL) DeleteReview(key uint) bool {

	//	Delete the item.
	cmdDelete := fmt.Sprintf("DELETE FROM reviews WHERE id = %d", key)
	if _, err := p.db.Exec(cmdDelete); err != nil {
		fmt.Printf("Error deleting item from db: %q\n", err)
		return false
	}
	fmt.Println("Deleted item from database")

	return true
}
func (p *ProviderPSQL) GetReviews() []models.Review {

	//	Read all rows.
	rows, err := p.db.Query("SELECT * FROM reviews")
	if err != nil {
		fmt.Printf("Error reading ticks: %q\n", err)
		return nil
	}

	//	Struct all the row data.
	reviews := make([]models.Review, 0)
	defer rows.Close()
	for rows.Next() {
		var id uint
		var kind models.ReviewType
		var name string
		var balance float32
		if err := rows.Scan(&id, &kind, &name, &balance); err != nil {
			fmt.Printf("Error scanning reviews: %q\n", err)
			return nil
		}

		reviews = append(reviews, models.Review{Key: id, Type: kind, Name: name, Balance: balance})
		// fmt.Printf("Read from DB: [id:%s] [Kind:%s] [Name:%s] [Balance:%s]\n", fmt.Sprintf("%d", id), kind, name, fmt.Sprintf("%.2f", balance))
	}

	return reviews
}

func (p *ProviderPSQL) GetSumAssets() (bool, float32) {

	//	Sum the Assets.
	var sumAssets float32
	cmdSumAssets := fmt.Sprintf("SELECT SUM(balance) FROM reviews WHERE type = '%s'",
		"Asset")
	rows, err := p.db.Query(cmdSumAssets)
	if err != nil {
		fmt.Printf("Error summing type-Asset reviews: %q\n", err)
		return false, 0.0
	}
	defer rows.Close()
	for rows.Next() {
		var amount float32
		if err := rows.Scan(&amount); err != nil {
			fmt.Printf("Error scanning reviews: %q\n", err)
			return false, sumAssets
		}
		sumAssets += amount
	}
	//fmt.Printf("Sum balance of assets in DB: [id:%.2f]\n", sumAssets)

	//
	return true, sumAssets

}
func (p *ProviderPSQL) GetSumLiabilities() (bool, float32) {

	//	Sum the Liabilities.
	var sumLiabilities float32
	cmdSumLiabilities := fmt.Sprintf("SELECT SUM(balance) FROM reviews WHERE type = '%s'",
		"Liability")
	rows, err := p.db.Query(cmdSumLiabilities)
	if err != nil {
		fmt.Printf("Error summing type-Liability reviews: %q\n", err)
		return false, 0.0
	}
	defer rows.Close()
	for rows.Next() {
		var amount float32
		if err := rows.Scan(&amount); err != nil {
			fmt.Printf("Error scanning reviews: %q\n", err)
			return false, sumLiabilities
		}
		sumLiabilities += amount
	}
	//fmt.Printf("Sum balance of liabilities in DB: [id:%.2f]\n", sumLiabilities)

	//
	return true, sumLiabilities
}
func (p *ProviderPSQL) GetLastReviewId() (bool, uint) {

	//	Find the largest key.
	var lastReviewId uint
	rows, err := p.db.Query("SELECT MAX(id) FROM reviews")
	if err != nil {
		fmt.Printf("Error finding latest review: %q\n", err)
		return false, 0
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&lastReviewId); err != nil {
			fmt.Printf("Error scanning reviews: %q\n", err)
			return false, 0
		}
	}
	//fmt.Printf("Latest review key in DB: [id:%d]\n", lastReviewId)

	//
	return true, lastReviewId
}

func (p *ProviderPSQL) Init() bool {

	//	Connect.
	//const connStr = "user=backend password=1234b dbname=backend sslmode=disable"
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("Error opening database:[%s] %q", connStr, err)
		return false
	}
	fmt.Printf("Successfully accessed database [%s]", connStr)
	p.db = db

	//	Make table if it doesn't exist yet.
	cmdCreateTable := fmt.Sprintf("CREATE TABLE IF NOT EXISTS reviews (%s, %s, %s, %s)",
		"id BIGSERIAL NOT NULL PRIMARY KEY",
		"type VARCHAR(100) NOT NULL",
		"name VARCHAR(100) NOT NULL",
		"balance REAL NOT NULL")
	if _, err := db.Exec(cmdCreateTable); err != nil {
		fmt.Printf("Error finding/creating database table: %q\n", err)
		return false
	}

	return true
}
