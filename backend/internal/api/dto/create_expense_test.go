package dto

import (
	"testing"

	"github.com/dKariakin/purser/internal/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateExpenseDto(t *testing.T) {
	t.Parallel()

	t.Run("Create expense request should be converted to a domain model when ToDomain is called",
		func(t *testing.T) {
			t.Parallel()

			req := CreateExpenseRequest{
				Price: 421,
				Title: " test Title",
			}
			expected := model.Expense{
				Name:  "test Title",
				Price: 421,
			}
			actual := req.ToDomain()

			assert.Equal(t, expected, actual)
		})

	t.Run("Create expense response should be created from a domain model when FromDomain is called",
		func(t *testing.T) {
			t.Parallel()

			domain := model.Expense{
				Id:    "1",
				Name:  "test",
				Price: 11,
			}
			expected := CreateExpenseResponse{
				CreateExpenseRequest: CreateExpenseRequest{
					Price: 11,
					Title: "test",
				},
				Id: "1",
			}
			result := ResponseFromDomain(domain)

			assert.Equal(t, expected, result)
		})
}
