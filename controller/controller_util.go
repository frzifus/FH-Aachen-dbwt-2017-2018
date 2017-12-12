package controllers

import (
	"github.com/gorilla/sessions"
	"net/http"
	"reflect"
)

const (
	cookieName = "session"
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

func readID(r *http.Request, store sessions.Store) (uint, error) {
	s, err := store.Get(r, cookieName)
	if err != nil {
		return 0, err
	}
	switch v := reflect.ValueOf(s.Values["id"]); v.Kind() {
	case reflect.Uint:
		return uint(v.Uint()), nil
	}
	return 0, err
}
