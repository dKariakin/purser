package error_codes

type ErrorCode int8

const (
	UnprocessedError ErrorCode = iota
	IncorrectContentType
	MalformedRequest
	ServiceTimeout
	InvalidPrice
	EmptyName
)

func (e ErrorCode) String() string {
	defaultErrCode := "UnprocessedError"
	errCodes := [...]string{
		defaultErrCode,
		"IncorrectContentType",
		"MalformedRequest",
		"ServiceTimeout",
		"InvalidPrice",
		"EmptyName",
	}

	if e < 0 || e > EmptyName {
		return defaultErrCode
	}

	return errCodes[e]
}
