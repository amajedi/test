package app

import (
	"net/http"
)

type AppError struct {
    Message string
    Code    int
}

type AppHandler func(http.ResponseWriter, *http.Request) *AppError

func (fn AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if err := fn(w, r); err != nil {
        http.Error(w, err.Message, err.Code)
    }
}