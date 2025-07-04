package errors

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"error,omitempty"`
}

func (e *AppError) Error() string {
	return e.Message
}

func New(code int, message string) error {
	return &AppError{Code: code, Message: message}
}
