package common

import (
	"errors"
	"net/http"
)

var (
	ErrNotFound = errors.New(http.StatusText(http.StatusNotFound))
	ErrInternal = errors.New(http.StatusText(http.StatusInternalServerError))
)
