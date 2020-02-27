package defs

// Err model
type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

// ErrorResponse model
type ErrorResponse struct {
	HTTPSC int
	Error  Err
}

// Error Defs
var (
	ErrorRequestBodyParseFailed = ErrorResponse{
		HTTPSC: 400,
		Error:  Err{Error: "Request body is not correct", ErrorCode: "001"},
	}

	ErrorNotAuthUser = ErrorResponse{
		HTTPSC: 401,
		Error:  Err{Error: "User authentication failed", ErrorCode: "002"},
	}
)
