package dto

type ErrorResponse struct {
	Errors []string `json:"errors,omitempty"`
}

func NewError(err string) ErrorResponse {
	return ErrorResponse{
		Errors: []string{err},
	}
}

func (e *ErrorResponse) Add(err string) ErrorResponse {
	if e.Errors == nil {
		return NewError(err)
	} else {
		e.Errors = append(e.Errors, err)
	}

	return *e
}
