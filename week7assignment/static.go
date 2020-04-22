package main

import (
	"net/http"
)

func Static() http.Handler {
	fs := http.FileServer(http.Dir(".//static"))
	return fs
}
