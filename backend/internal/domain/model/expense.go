package model

import "time"

type Expense struct {
	Id        string
	Name      string
	Price     float32
	CreatedAt time.Time
}
