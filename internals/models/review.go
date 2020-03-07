package models

import "fmt"

type ReviewType string

const (
	ReviewTypeAsset     = "Asset"
	ReviewTypeLiability = "Liability"
)

type Review struct {
	Key     uint
	Type    ReviewType
	Name    string
	Balance float32
}

func (r *Review) GetString() string {
	return fmt.Sprintf("review - [Key:%d] [Type:%s] [Name:%s] [Balance:%f]",
		r.Key, r.Type, r.Name, r.Balance)
}
func (r *Review) CopyFrom(other *Review) {
	r.Type = other.Type
	r.Name = other.Name
	r.Balance = other.Balance
}

func (r *Review) SetTypeAsset() {
	r.Type = ReviewTypeAsset
}
func (r *Review) SetTypeLiability() {
	r.Type = ReviewTypeLiability
}
func (r *Review) IsAsset() bool {
	return r.Type == ReviewTypeAsset
}
func (r *Review) IsLiability() bool {
	return r.Type == ReviewTypeLiability
}
