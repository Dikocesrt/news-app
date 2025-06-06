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
var ErrUnauthorized = errors.New("unauthorized")
var ErrCustomPageAlreadyExists = errors.New("custom page already exists")
var ErrInvalidCustomPageID = errors.New("invalid custom page id")
var ErrInvalidCustomURL = errors.New("custom url must not contain spaces")
var ErrTagAlreadyExists = errors.New("tag already exists")

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
		case ErrUnauthorized:
			return http.StatusUnauthorized
		case ErrCustomPageAlreadyExists:
			return http.StatusConflict
		case ErrInvalidCustomPageID:
			return http.StatusBadRequest
		case ErrInvalidCustomURL:
			return http.StatusBadRequest
		case ErrTagAlreadyExists:
			return http.StatusConflict
		default:
			return http.StatusInternalServerError
	}
}