package domain

import "errors"

// Universal errors
var (
	ErrNotFound  = errors.New("not found")
	ErrForbidden = errors.New("forbidden")
	ErrNotUnique = errors.New("not unique")
	ErrMismatch  = errors.New("mismatch")
)

// Special domain errors
var (
	ErrUnauthorized       = errors.New("unauthorized")
	ErrBadFormat          = errors.New("bad format")
	ErrNotConfirmed       = errors.New("identity is not confirmed")
	ErrUnknownIdent       = errors.New("unknown identity")
	ErrEmailPassNotFound  = errors.New("email-pass identity not found")
	ErrIdentityDuplicated = errors.New("identity duplicated")
)
