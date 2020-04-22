package main

import (
	"log"
	"net/http"
)

func main() {
	fs1 := Static()
	fs2 := FileServer()

	http.Handle("/", http.StripPrefix("/", fs1))
	http.Handle("/fileserver", http.StripPrefix("/fileserver", fs2))

	log.Println("Listening on 3000")
	http.ListenAndServe(":3000", nil)

}
