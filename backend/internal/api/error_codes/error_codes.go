package error_codes

type ErrorCode int8

const (
	InvalidPrice ErrorCode = iota
	EmptyName
)

func (e ErrorCode) String() string {
	errCodes := [...]string{
		"InvalidPrice",
		"EmptyName",
	}

	if e < 0 || e > EmptyName {
		return "UnprocessedError"
	}

	return errCodes[e]
}
