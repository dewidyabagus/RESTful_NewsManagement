package business

import "errors"

var (
	ErrDataNotSpec = errors.New("data not spec")

	ErrDataConflict = errors.New("data conflict")

	ErrUnauthorized = errors.New("unauthorized")

	ErrDataNotFound = errors.New("data not found")

	ErrBadRequest = errors.New("bad request")

	ErrHasBeenPublished = errors.New("has been published")
)
