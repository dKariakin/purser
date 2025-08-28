package dto

import (
	"strings"

	"github.com/dKariakin/purser/internal/domain/model"
)

type CreateExpenseRequest struct {
	Price float32 `json:"price"`
	Title string  `json:"title"`
}

type CreateExpenseResponse struct {
	CreateExpenseRequest
	Id string `json:"id"`
}

func (req *CreateExpenseRequest) ToDomain() model.Expense {
	return model.Expense{
		Name:  strings.TrimSpace(req.Title),
		Price: req.Price,
	}
}

func ResponseFromDomain(expense model.Expense) CreateExpenseResponse {
	return CreateExpenseResponse{
		CreateExpenseRequest: CreateExpenseRequest{
			Price: expense.Price,
			Title: expense.Name,
		},
		Id: expense.Id,
	}
}
