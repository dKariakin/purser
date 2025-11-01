package db

import "github.com/dKariakin/purser/internal/domain/model"

type DbClient interface {
	CreateExpense() (model.Expense, error)
}
