package models

import "time"

type Test struct {
	Id   int64  `json:"id"`
	Test string `json:"test"`
	TestInt int `json:"testInt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
