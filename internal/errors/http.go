package errors

import "net/http"

// StatusFor maps a sentinel error to the HTTP status code the API should return for it.
func StatusFor(err error) int {
	switch err {
	case ErrParentNotFound, ErrTeacherNotFound:
		return http.StatusNotFound
	case ErrInvalidRequest:
		return http.StatusBadRequest
	case ErrRateLimited:
		return http.StatusTooManyRequests
	default:
		return http.StatusInternalServerError
	}
}

// Code returns the snake_case wire code for a sentinel error, falling back to internal_error.
func Code(err error) string {
	switch err {
	case ErrParentNotFound, ErrTeacherNotFound, ErrRateLimited, ErrInvalidRequest:
		return err.Error()
	default:
		return ErrInternal.Error()
	}
}
