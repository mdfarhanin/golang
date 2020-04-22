package main

import (
	"net/http"
)

func FileServer() http.Handler {
	dir := http.Dir("C://")
	fs := http.FileServer(dir)
	return fs
}
