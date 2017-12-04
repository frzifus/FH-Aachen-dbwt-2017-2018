package controllers

import (
	"github.com/gorilla/sessions"
	"net/http"
	"reflect"
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

func role(r *http.Request, store sessions.Store) string {
	s, err := store.Get(r, "SessionOtherCookie")
	if err != nil {
		return ""
	}
	if _, ok := s.Values["role"]; ok {
		r := reflect.ValueOf(s.Values["role"]).String()
		return r
	}
	return ""
}
