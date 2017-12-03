package controllers

import (
	"github.com/gorilla/sessions"
	"net/http"
)

func signedIn(r *http.Request, store sessions.Store) bool {
	s, err := store.Get(r, "SomeOtherCookie")
	if err != nil {
		return false
	}
	if s.Values["active"] == true {
		return true
	}
	return false
}
