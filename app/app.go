package app

import (
	"net/http"
)

type AppError struct {
    Message string
    Code    int
}

type ApiContext struct {
	Params map[string]string
}

type AppHandler func(http.ResponseWriter, *http.Request, *ApiContext) *AppError

func (fn AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := ApiContext{
		make(map[string]string)}
    if err := fn(w, r, &context); err != nil {
        http.Error(w, err.Message, err.Code)
    }
}