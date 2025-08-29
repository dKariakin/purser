package create

import (
	"fmt"
	"testing"
	"time"

	"github.com/dKariakin/purser/internal/api/dto"
	"github.com/dKariakin/purser/internal/api/error_codes"
	"github.com/stretchr/testify/assert"
)

func TestNewExpenseHandler(t *testing.T) {
	t.Parallel()

	t.Run("Validation of a correct request should not rise any errors", func(t *testing.T) {
		t.Parallel()

		handler := New(nil, 1*time.Second)
		req := dto.CreateExpenseRequest{
			Price: 42,
			Title: " title",
		}
		result := handler.validateRequest(req)

		assert.Empty(t, result)
	})

	t.Run("When request has multiple errors, all of them should be returned from validator", func(t *testing.T) {
		t.Parallel()

		handler := New(nil, 1*time.Second)
		req := dto.CreateExpenseRequest{
			Price: 0,
			Title: " ",
		}
		expected := []error{
			fmt.Errorf("%s", error_codes.InvalidPrice),
			fmt.Errorf("%s", error_codes.EmptyName),
		}
		result := handler.validateRequest(req)

		assert.Equal(t, expected, result)
	})
}
