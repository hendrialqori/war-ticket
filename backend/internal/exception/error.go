package exception

type AppError struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     any    `json:"errors,omitempty"`
}

func (e *AppError) Error() string {
	return e.Message
}

func New(code int, message string) *AppError {
	return &AppError{
		Success: false,
		Code:    code,
		Message: message,
	}
}

func NewWithError(code int, message string, err any) *AppError {
	return &AppError{
		Success: false,
		Code:    code,
		Message: message,
		Err:     err,
	}
}
