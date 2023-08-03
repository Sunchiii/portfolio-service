package utils

import "net/http"

// ErrorResponse represents a JSON error response.
type ErrorResponse struct {
    Message string `json:"message"`
}

// NewErrorResponse creates a new ErrorResponse instance.
func NewErrorResponse(message string) *ErrorResponse {
    return &ErrorResponse{
        Message: message,
    }
}

// HTTPError represents an HTTP error.
type HTTPError struct {
    Status  int
    Message string
}

// Error returns the error message for an HTTPError instance.
func (e *HTTPError) Error() string {
    return e.Message
}

// NewHTTPError creates a new HTTPError instance.
func NewHTTPError(status int, message string) *HTTPError {
    return &HTTPError{
        Status:  status,
        Message: message,
    }
}

// BadRequestError creates a new HTTPError instance with status code 400.
func BadRequestError(message string) *HTTPError {
    return NewHTTPError(http.StatusBadRequest, message)
}

// UnauthorizedError creates a new HTTPError instance with status code 401.
func UnauthorizedError(message string) *HTTPError {
    return NewHTTPError(http.StatusUnauthorized, message)
}

// NotFoundError creates a new HTTPError instance with status code 404.
func NotFoundError(message string) *HTTPError {
    return NewHTTPError(http.StatusNotFound, message)
}

// InternalServerError creates a new HTTPError instance with status code 500.
func InternalServerError(message string) *HTTPError {
    return NewHTTPError(http.StatusInternalServerError, message)
}
