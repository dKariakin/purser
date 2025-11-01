package dto

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/dKariakin/purser/internal/domain/model"
)

type Expense struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func FromDomain(e model.Expense) Expense {
	return Expense{
		Id:    e.Id,
		Name:  e.Name,
		Price: e.Price,
	}
}

func (e Expense) Reader() io.Reader {
	j, err := json.Marshal(e)
	if err != nil {
		return nil
	}

	return strings.NewReader(string(j))
}
