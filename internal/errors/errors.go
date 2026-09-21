// Package errors defines sentinel errors shared across services and handlers.
// Codes are snake_case so they can be surfaced directly in API error responses.
package errors

import "errors"

var (
	ErrParentNotFound  = errors.New("parent_not_found")
	ErrTeacherNotFound = errors.New("teacher_not_found")
	ErrRateLimited     = errors.New("rate_limited")
	ErrInvalidRequest  = errors.New("invalid_request")
	ErrInternal        = errors.New("internal_error")
)
