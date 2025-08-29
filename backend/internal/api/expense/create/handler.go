package create

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dKariakin/purser/internal/api/dto"
	"github.com/dKariakin/purser/internal/api/error_codes"
	"github.com/labstack/echo/v4"
)

type NewExpenseHandler struct {
	service Service
	timeout time.Duration
}

// New creates a new instance of NewExpenseHandler
func New(svc Service, t time.Duration) *NewExpenseHandler {
	return &NewExpenseHandler{
		service: svc,
		timeout: t,
	}
}

// CreateExpense handles request of creating a new expense
func (handler *NewExpenseHandler) CreateExpense(ctx echo.Context) error {
	request := dto.CreateExpenseRequest{}
	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, `{"error": "MalformedRequest"}`) // Add error handler for such cases
	}

	if err := handler.validateRequest(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, `{}`)
	}

	internalCtx, cancel := context.WithTimeout(context.Background(), handler.timeout)
	defer cancel()

	createdExpense, err := handler.service.CreateExpense(internalCtx, request.ToDomain())
	if err != nil {
		return err // Add error handler for such cases
	}

	return ctx.JSON(http.StatusCreated, dto.ResponseFromDomain(createdExpense))
}

// validateRequest validates body of the request, returning information regarding all errors
func (handler *NewExpenseHandler) validateRequest(request dto.CreateExpenseRequest) []error {
	var res []error

	if request.Price <= 0 {
		res = append(res, fmt.Errorf("%s", error_codes.InvalidPrice))
	}
	if len(strings.TrimSpace(request.Title)) == 0 {
		res = append(res, fmt.Errorf("%s", error_codes.EmptyName))
	}

	return res
}
