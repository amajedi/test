package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/amajedi/test/app"
	"github.com/amajedi/test/validation"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, c *app.ApiContext) *app.AppError {
	
	//
	// Perform request validations
	//
	if err := validation.ValidateIndex(r, c); err != nil {
		return err
	}

	//
	// Fetch target bytes of remote resource
	//
	req, reqErr := http.NewRequest(http.MethodGet, c.Params["source"], nil)
	if reqErr != nil {
		fmt.Println("Unable to construct GET request, context:")
		fmt.Println(c)
		return &app.AppError{"Internal Server Error", 500}
	}

	//
	// Issue remote request
	//
	rangeHeaderValue := fmt.Sprintf("bytes=%s", c.Params["range"])
	rangeHeaderKey := http.CanonicalHeaderKey("range")
	req.Header.Set(rangeHeaderKey, rangeHeaderValue)
	resp, respErr := app.Client.Do(req)
	if (resp.StatusCode != http.StatusPartialContent && 
		resp.StatusCode != http.StatusOK) && respErr != nil {
		return &app.AppError{"Unable to fetch data from source provided.", 422}
	}

	//
	// Read body
	//
	defer resp.Body.Close()
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		fmt.Println("Problem reading source response, context:")
		fmt.Println(c)
		fmt.Println("Error:")
		fmt.Println(readErr)
		return &app.AppError{"Problem reading source response", 422}
	}

	//
	// Respond
	//
	delete(resp.Header, http.CanonicalHeaderKey("accept-ranges"))
	for name, value := range resp.Header {
	    w.Header().Set(name, value[0])
	}
	w.Write(body)

	return nil
}



