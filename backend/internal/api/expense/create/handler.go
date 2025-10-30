package create

import (
	"context"
	"errors"
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
	request := &dto.CreateExpenseRequest{}
	if ctx.Request().Header.Get("Content-Type") != "application/json" {
		return ctx.JSON(http.StatusBadRequest, dto.NewError(error_codes.IncorrectContentType.String()))
	}

	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.NewError(error_codes.MalformedRequest.String()))
	}

	if err := handler.validateRequest(request); err.Errors != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	internalCtx, cancel := context.WithTimeout(context.Background(), handler.timeout)
	defer cancel()

	createdExpense, err := handler.service.CreateExpense(internalCtx, request.ToDomain())
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return ctx.JSON(http.StatusRequestTimeout, dto.NewError(error_codes.ServiceTimeout.String()))
		}

		return ctx.JSON(http.StatusInternalServerError, dto.NewError(error_codes.UnprocessedError.String()))
	}

	return ctx.JSON(http.StatusCreated, dto.ResponseFromDomain(createdExpense))
}

// validateRequest validates body of the request, returning information regarding all errors
func (handler *NewExpenseHandler) validateRequest(request *dto.CreateExpenseRequest) dto.ErrorResponse {
	var res dto.ErrorResponse

	if request.Price <= 0 {
		res = res.Add(error_codes.InvalidPrice.String())
	}
	if len(strings.TrimSpace(request.Title)) == 0 {
		res = res.Add(error_codes.EmptyName.String())
	}

	return res
}
