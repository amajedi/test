package validation

import (
	"fmt"
	"strings"
	"net/http"

	"github.com/amajedi/test/app"
)

func ValidateSource(r *http.Request, apiContext *app.ApiContext) *app.AppError {
	//
	// Source URL parameter required, 's'
	//
	sourceSlice, ok := r.URL.Query()["s"]; 
	if !ok {
		return &app.AppError{"Source parameter (s), required.", 422}
	}

	//
	// Must support range requests
	//
	source := sourceSlice[0]
	sourceHeadRes, sourceHeadErr := app.Client.Head(source)
	if sourceHeadErr != nil || (sourceHeadRes.StatusCode / 100) != 2 {
		return &app.AppError{
	    	"Source parameter (s), unable to validate, " +
	    	"resource not reachable or does not support HTTP HEAD requests.", 
	    	422}
	}
	cHeaderKey := http.CanonicalHeaderKey("accept-ranges")
	acceptRangesValue, acceptRangesOk := sourceHeadRes.Header[cHeaderKey]
	if !acceptRangesOk || strings.ToLower(acceptRangesValue[0]) != "bytes" {
		return &app.AppError{
	    	fmt.Sprintf("Source parameter (s), provided resource does not indicate " +
	    		"support for byte range requests via the '%s' response header.", 
	    		cHeaderKey), 
	    	422}
	}

	//
	// Update context
	//
	apiContext.Params["source"] = source
	
	return nil
}