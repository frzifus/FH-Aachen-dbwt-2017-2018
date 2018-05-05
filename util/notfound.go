package util

import (
	"log"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Oops! 404"))
	if err != nil {
		log.Fatal(err)
	}
}
