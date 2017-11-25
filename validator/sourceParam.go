package validator

import (
	"net/http"

	"github.com/amajedi/test/app"
)

func ValidateSource(r *http.Request) *app.AppError {
	// Source URL parameter required, 's'
	sourceSlice, ok := r.URL.Query()["s"]; 
	if !ok {
		return &app.AppError{"Source parameter (s) required.", 422}
	}

	// source url must support range requests
	// todo

	return nil
}