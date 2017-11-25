package validator

import (
	"net/http"

	"github.com/amajedi/test/app"
)

func ValidateIndex(r *http.Request) *app.AppError {
	//
	// source URL query parameter validation
	//
	if err := ValidateSource(r); err != nil {
		return err
	}
	//
	// range URL query parameter validation
	//
	if err := ValidateRange(r); err != nil {
		return err
	}

	return nil
}