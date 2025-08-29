package create

import (
	"context"

	"github.com/dKariakin/purser/internal/domain/model"
)

type Service interface {
	CreateExpense(context.Context, model.Expense) (model.Expense, error)
}

//go:generate mockgen -destination=mocks/service.go -package=mocks . Service
