package controllers

import (
	"github.com/gorilla/sessions"
	"net/http"
	"reflect"
)

var (
	cookieName = "SomeOtherCookie"
)

func signedIn(r *http.Request, store sessions.Store) bool {
	s, err := store.Get(r, cookieName)
	if err != nil {
		return false
	}
	if s.Values["active"] == true {
		return true
	}
	return false
}

func role(r *http.Request, store sessions.Store) string {
	s, err := store.Get(r, cookieName)
	if err != nil {
		return ""
	}
	switch v := reflect.ValueOf(s.Values["role"]); v.Kind() {
	case reflect.String:
		return v.String()
	}
	return ""
}
