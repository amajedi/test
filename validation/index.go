package validation

import (
	"net/http"

	"github.com/amajedi/test/app"
)

func ValidateIndex(r *http.Request, c *app.ApiContext) *app.AppError {

	//
	// Validate Path
	//
	if r.URL.Path != "/" {
		return &app.AppError{"Not Found", 404}
	}

	//
	// Validate Query Parameters
	//
	if err := ValidateSource(r, c); err != nil {
		return err
	}
	if err := ValidateRange(r, c); err != nil {
		return err
	}

	return nil
}