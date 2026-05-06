package apperror

import "net/http"

type AppError struct {
	StatusCode int
	Message    string
	Err        error
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func BadRequestError(message string, err error) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		Message:    message,
		Err:        err,
	}
}

func UnauthorizedError(message string, err error) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		Message:    message,
		Err:        err,
	}
}

func ForbiddenError(message string, err error) *AppError {
	return &AppError{
		StatusCode: http.StatusForbidden,
		Message:    message,
		Err:        err,
	}
}

func NotFoundError(message string, err error) *AppError {
	return &AppError{
		StatusCode: http.StatusNotFound,
		Message:    message,
		Err:        err,
	}
}

func ConflictError(message string, err error) *AppError {
	return &AppError{
		StatusCode: http.StatusConflict,
		Message:    message,
		Err:        err,
	}
}

func UnprocessableEntityError(message string, err error) *AppError {
	return &AppError{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    message,
		Err:        err,
	}
}

func TooManyRequestsError(message string, err error) *AppError {
	return &AppError{
		StatusCode: http.StatusTooManyRequests,
		Message:    message,
		Err:        err,
	}
}

func InternalError(err error) *AppError {
	return &AppError{
		StatusCode: http.StatusInternalServerError,
		Message:    "internal server error",
		Err:        err,
	}
}

func InternalServerError(message string, err error) *AppError {
	return &AppError{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
		Err:        err,
	}
}

func ServiceUnavailableError(message string, err error) *AppError {
	return &AppError{
		StatusCode: http.StatusServiceUnavailable,
		Message:    message,
		Err:        err,
	}
}
