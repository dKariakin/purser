package create

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/dKariakin/purser/internal/api/dto"
	"github.com/dKariakin/purser/internal/api/error_codes"
	"github.com/dKariakin/purser/internal/api/expense/create/mocks"
	"github.com/dKariakin/purser/internal/domain/model"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNewExpenseHandler(t *testing.T) {
	t.Parallel()

	e := echo.New()

	t.Run("When a new expense is successfully created, its details should be returned in API response",
		func(t *testing.T) {
			t.Parallel()

			svcResult := model.Expense{
				Id:        "1",
				Name:      "My expense",
				Price:     42,
			}
			expected := fmt.Sprintln(`{"id":"1","price":42,"title":"My expense"}`)
			r := strings.NewReader(`{"price":42, "title":"My expense"}`)

			req := httptest.NewRequest(http.MethodPost, "/", r)
			req.Header.Add("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			ctrl := gomock.NewController(t)
			mock := mocks.NewMockService(ctrl)
			handler := New(mock, 1)

			mock.EXPECT().CreateExpense(gomock.Any(), gomock.Any()).Return(svcResult, nil).Times(1)

			handler.CreateExpense(ctx)

			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, expected, rec.Body.String())
		})

	t.Run("When request has an incorrect body, a generic error should be returned", func(t *testing.T) {
		t.Parallel()

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{`))
		req.Header.Add("Content-Type", "application/json")

		ctx := e.NewContext(req, rec)

		handler := New(nil, 1)
		handler.CreateExpense(ctx)

		expected := fmt.Sprintln(`{"errors":["MalformedRequest"]}`)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, expected, rec.Body.String())
	})

	t.Run("When Content-Type header is not set, handler should result to 400 error", func(t *testing.T) {
		t.Parallel()

		expected := fmt.Sprintln(`{"errors":["IncorrectContentType"]}`)
		body := strings.NewReader(`{"price":1,"title":"nope"}`)
		req := httptest.NewRequest(http.MethodPost, "/", body)
		rec := httptest.NewRecorder()

		ctx := e.NewContext(req, rec)

		handler := New(nil, 1)

		handler.CreateExpense(ctx)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, expected, rec.Body.String())
	})

	t.Run("When request times out, handler should result to 408 error", func(t *testing.T) {
		t.Parallel()

		body := strings.NewReader(`{"price":1,"title":"timeout"}`)
		expected := fmt.Sprintln(`{"errors":["ServiceTimeout"]}`)
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/", body)
		req.Header.Add("Content-Type", "application/json")

		ctx := e.NewContext(req, rec)

		ctrl := gomock.NewController(t)
		svc := mocks.NewMockService(ctrl)
		handler := New(svc, 1)

		svc.EXPECT().CreateExpense(gomock.Any(), gomock.Any()).Return(model.Expense{}, context.DeadlineExceeded).Times(1)

		handler.CreateExpense(ctx)

		assert.Equal(t, http.StatusRequestTimeout, rec.Code)
		assert.Equal(t, expected, rec.Body.String())
	})

	t.Run("When expense service returns an error, it results to 500 error code", func(t *testing.T) {
		t.Parallel()

		body := strings.NewReader(`{"price":1,"title":"boom"}`)
		expected := fmt.Sprintln(`{"errors":["UnprocessedError"]}`)

		ctrl := gomock.NewController(t)
		svc := mocks.NewMockService(ctrl)
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/", body)
		req.Header.Add("Content-Type", "application/json")

		svc.EXPECT().CreateExpense(gomock.Any(), gomock.Any()).Return(model.Expense{}, errors.New("boom")).Times(1)

		ctx := e.NewContext(req, rec)
		handler := New(svc, 1)

		handler.CreateExpense(ctx)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, expected, rec.Body.String())
	})

	t.Run("Validation of a correct request should not rise any errors", func(t *testing.T) {
		t.Parallel()

		handler := New(nil, 1*time.Second)
		req := &dto.CreateExpenseRequest{
			Price: 42,
			Title: " title",
		}
		result := handler.validateRequest(req)

		assert.Empty(t, result)
	})

	t.Run("When request has multiple errors, all of them should be returned from validator", func(t *testing.T) {
		t.Parallel()

		handler := New(nil, 1*time.Second)
		req := &dto.CreateExpenseRequest{
			Price: 0,
			Title: " ",
		}
		expected := dto.NewError(error_codes.InvalidPrice.String())
		expected.Add(error_codes.EmptyName.String())

		result := handler.validateRequest(req)

		assert.Equal(t, expected, result)
	})
}
