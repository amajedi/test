package validation

import (
	"strings"
	"strconv"
	"net/http"

	"github.com/amajedi/test/app"
)

func ValidateRange(r *http.Request, apiContext *app.ApiContext) *app.AppError {

	query := r.URL.Query()

	//
	// range URL parameter required, 'r'
	//
	rangeSlice, rangeOk := query["range"];
	if !rangeOk {
		return &app.AppError{"Range parameter (range) required.", 422}
	}

	//
	// range URL parameter of the form, "start-end" where end optional and both integers
	//
	rangeTokens := strings.Split(rangeSlice[0], "-")
	if len(rangeTokens) != 2 {
		return &app.AppError{
			"Range parameter (range) not valid, not of the form `start-end`, end-byte optional.", 
			422}
	}

	//
	// validate start-byte, ensure uint
	//
	start, startErr := strconv.ParseUint(rangeTokens[0], 10, 64)
	if startErr != nil {
		return &app.AppError{"Range parameter (range) not valid, start-byte value.", 422}
	}

	//
	// validate end-byte, ensure uint, and greater than start-byte
	//
	if rangeTokens[1] != "" {
		end, endErr := strconv.ParseUint(rangeTokens[1], 10, 64)
		if endErr != nil {
			return &app.AppError{"Range parameter (range) not valid, end-byte value.", 422}
		}
		if end < start {
			return &app.AppError{
				"Range parameter (range) not valid, end-byte less than start-byte.", 422}
		}
	}
	
	//
	// Update context
	//
	apiContext.Params["range"] = rangeSlice[0]

	return nil
}