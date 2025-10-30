package model

import "time"

type Expense struct {
	Id        string
	Name      string
	Price     int
	CreatedAt time.Time
}
