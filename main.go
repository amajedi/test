package main

import (
	"net/http"

	"github.com/amajedi/test/app"
	"github.com/amajedi/test/controller"
)

func main() {
	http.Handle("/", app.AppHandler(controller.IndexHandler))
	http.ListenAndServe(":8000", nil)
}