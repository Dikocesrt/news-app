package utils

import (
	"errors"
	"net/http"
)

var ErrEmptyField = errors.New("field cannot be empty")
var ErrUsernameAlreadyRegistered = errors.New("username already registered")
var ErrInvalidToken = errors.New("invalid token")
var ErrInvalidCredentials = errors.New("username or password is incorrect")
var ErrInvalidCategoryID = errors.New("invalid category id")
var ErrCategoryAlreadyExists = errors.New("category already exists")
var ErrInvalidNewsID = errors.New("invalid news id")

func ConvertErrorCode(err error) int {
	switch err {
		case ErrEmptyField:
			return http.StatusBadRequest
		case ErrUsernameAlreadyRegistered:
			return http.StatusConflict
		case ErrInvalidToken:
			return http.StatusUnauthorized
		case ErrInvalidCredentials:
			return http.StatusUnauthorized
		case ErrInvalidCategoryID:
			return http.StatusBadRequest
		case ErrCategoryAlreadyExists:
			return http.StatusConflict
		case ErrInvalidNewsID:
			return http.StatusBadRequest
		default:
			return http.StatusInternalServerError
	}
}