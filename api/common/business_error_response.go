package common

import (
	"RESTful/business"
	"net/http"

	"gorm.io/gorm"
)

var (
	ErrDataNotSpec  = "400"
	ErrUnauthorized = "401"
	ErrDataNotFound = "404"
	ErrDataConflict = "409"

	ErrInternalServer = "500"
)

type BusinessErrorResponseSpec struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewBusinessErrorResponse(err error) (int, *BusinessErrorResponseSpec) {
	switch err {
	default:
		return errResponseInternalServerError()

	case business.ErrDataConflict:
		return errResponseDataConflict(err.Error())

	case business.ErrDataNotSpec:
		return errResponseDataNotSpec(err.Error())

	case business.ErrUnauthorized:
		return errResponseUnauthorized(err.Error())

	case gorm.ErrRecordNotFound:
		return errResponseDataNotFound()

	case business.ErrDataNotFound:
		return errResponseDataNotFound()

	case business.ErrBadRequest:
		return errResponseBadRequest(err.Error())
	}
}

func errResponseInternalServerError() (int, *BusinessErrorResponseSpec) {
	return http.StatusInternalServerError, &BusinessErrorResponseSpec{
		Code:    ErrInternalServer,
		Message: "internal server error",
	}
}

func errResponseDataConflict(message string) (int, *BusinessErrorResponseSpec) {
	return http.StatusConflict, &BusinessErrorResponseSpec{
		Code:    ErrDataConflict,
		Message: message,
	}
}

func errResponseDataNotSpec(message string) (int, *BusinessErrorResponseSpec) {
	return http.StatusBadRequest, &BusinessErrorResponseSpec{
		Code:    ErrDataNotSpec,
		Message: message,
	}
}

func errResponseUnauthorized(message string) (int, *BusinessErrorResponseSpec) {
	return http.StatusUnauthorized, &BusinessErrorResponseSpec{
		Code:    ErrUnauthorized,
		Message: message,
	}
}

func errResponseDataNotFound() (int, *BusinessErrorResponseSpec) {
	return http.StatusNotFound, &BusinessErrorResponseSpec{
		Code:    ErrDataNotFound,
		Message: "data not found",
	}
}

func errResponseBadRequest(message string) (int, *BusinessErrorResponseSpec) {
	return http.StatusBadRequest, &BusinessErrorResponseSpec{
		Code:    ErrBadRequest,
		Message: message,
	}
}
