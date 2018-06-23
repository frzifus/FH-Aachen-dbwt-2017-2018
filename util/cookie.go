package util

import (
	"net/http"
	"reflect"

	"git.klimlive.de/frzifus/dbwt/model"
	"github.com/gernest/utron/base"
)

const (
	cookieName = "session"
)

func NewSession(ctx *base.Context, u *model.User) {
	service := model.NewService(ctx)
	session, _ := ctx.NewSession(cookieName)
	session.Values["username"] = u.Loginname
	session.Values["id"] = u.ID
	session.Values["role"] = service.GetRole(u)
	session.Values["active"] = true
	session.Options.Path = "/"
	session.Options.MaxAge = 10 * 24 * 3600
	_ = session.Save(ctx.Request(), ctx.Response())
}

func DeleteSession(ctx *base.Context) {
	r := ctx.Request()
	session, _ := ctx.SessionStore.Get(r, cookieName)
	session.Options.MaxAge = -1
	_ = session.Save(r, ctx.Response())
	ctx.Redirect(r.Header.Get("Referer"), http.StatusFound)
}

func IsSignedIn() func(*base.Context) error {
	return func(ctx *base.Context) error {
		s, err := ctx.SessionStore.Get(ctx.Request(), cookieName)
		if err != nil {
			ctx.Data["signedIn"] = false
		}
		if s.Values["active"] == true {
			ctx.Data["signedIn"] = true
		}
		return nil
	}
}

func Role(ctx *base.Context) string {
	s, err := ctx.SessionStore.Get(ctx.Request(), cookieName)
	if err != nil {
		return ""
	}
	switch v := reflect.ValueOf(s.Values["role"]); v.Kind() {
	case reflect.String:
		return v.String()
	}
	return ""
}

func ReadID(ctx *base.Context) (uint, error) {
	s, err := ctx.SessionStore.Get(ctx.Request(), cookieName)
	if err != nil {
		return 0, err
	}
	switch v := reflect.ValueOf(s.Values["id"]); v.Kind() {
	case reflect.Uint:
		return uint(v.Uint()), nil
	}
	return 0, err
}
