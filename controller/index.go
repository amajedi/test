package controller

import (
	"net/http"

	"github.com/amajedi/test/app"
	"github.com/amajedi/test/validator"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) *app.AppError {
	if err := validator.ValidateIndex(r); err != nil {
		return err
	}
	return nil
}