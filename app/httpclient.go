package app

import (
	"net/http"
	"time"
)

var (
	Client *http.Client
)

func init() {
	Client = &http.Client{
		Timeout: time.Second * 30}
}
