package dto

import (
	"strings"
	"testing"

	"github.com/dKariakin/purser/internal/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestExpenseDto(t *testing.T) {
	t.Parallel()

	t.Run("When a new DB expense object is created from a domain model, createdAt should not be set", func(t *testing.T) {
		t.Parallel()

		source := model.Expense{
			Id:    "1",
			Name:  "test",
			Price: 42,
		}
		expected := Expense{
			Id:    "1",
			Name:  "test",
			Price: 42,
		}

		actual := FromDomain(source)

		assert.Equal(t, expected, actual)
	})

	t.Run("When an expense DTO is converted to reader, a correct object should be returned", func(t *testing.T) {
		t.Parallel()

		source := Expense{
			Id:    "42",
			Name:  " test expense",
			Price: 42,
		}
		expected := strings.NewReader(`{"id":"42","name":" test expense","price":42}`)
		result := source.Reader()

		assert.Equal(t, expected, result)
	})
}
