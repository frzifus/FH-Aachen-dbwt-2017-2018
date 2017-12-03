package controllers

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/frzifus/dbwt/model"
	"github.com/gernest/utron/controller"
	"net/http"
)


type login struct {
	controller.BaseController
	Routes []string
}

// NewLogin -
func NewLogin() controller.Controller {
	return &login{
		Routes: []string{
			"get;/SignUp;SignUp",
			"get;/SignIn;SignIn",
			"get;/SignOff;SignOff",
			"post;/Success;Success",
			"post;/Register;Register",
		},
	}
}

func (l *login) encryptPassword(password string) string {
	h := sha256.Sum256([]byte(password))
	return "{SHA256}" + base64.StdEncoding.EncodeToString(h[:])
}

func (l *login) SignIn() {
	l.Ctx.Data["signdIn"] = signdIn(l.Ctx.Request(), l.Ctx.SessionStore)
	if len(l.Ctx.Request().URL.Query().Get("error")) > 0 {
		l.Ctx.Data["error"] = true
	} else {
		l.Ctx.Data["error"] = false
	}
	l.Ctx.Template = "login/signin"
	l.HTML(http.StatusOK)
}

func (l *login) Success() {
	r := l.Ctx.Request()
	ln := r.FormValue("login_name")
	pwd := r.FormValue("password")
	pwd64 := l.encryptPassword(pwd)
	u := &model.User{}

	if err := l.Ctx.DB.Where("loginname = ? AND hash = ?",
		ln, pwd64).Find(&u).Error; err != nil {

		l.Ctx.Redirect("/SignIn?error=1", http.StatusFound)
	}

	session, _ := l.Ctx.NewSession("SomeOtherCookie")
	session.Values["username"] = u.Loginname
	session.Values["Role"] = "Monkey"
	session.Values["active"] = true
	session.Options.Path = "/"
	session.Options.MaxAge = 10 * 24 * 3600
	_ = session.Save(l.Ctx.Request(), l.Ctx.Response())

	l.Ctx.Data["user"] = u
	l.Ctx.Data["signdIn"] = true
	l.Ctx.Template = "login/success"
	l.HTML(http.StatusOK)
}

func (l *login) SignOff() {
	session, err := l.Ctx.SessionStore.Get(l.Ctx.Request(), "SomeOtherCookie")
	if err != nil {
		l.Ctx.Redirect("/", http.StatusFound)
	}
	session.Options.MaxAge = -1
	_ = session.Save(l.Ctx.Request(), l.Ctx.Response())
	l.Ctx.Redirect("/", http.StatusFound)
}

func (l *login) newSession(name string) error {
	session, _ := l.Ctx.NewSession(name)
	return l.Ctx.SaveSession(session)
}

func (l *login) Register() {
	r := l.Ctx.Request()
	newUser := &model.User{
		Active:    true,
		Firstname: r.FormValue("first_name"),
		Lastname:  r.FormValue("last_name"),
		Mail:      r.FormValue("email"),
		Loginname: r.FormValue("display_name"),
		Algo:      "sha256",
		Hash:      l.encryptPassword(r.FormValue("password")),
	}

	if err := l.Ctx.DB.Create(newUser).Error; err != nil {
		l.Ctx.DB.Rollback()
		fmt.Printf("Could not create user: %s", err)
	}

	l.Ctx.Redirect("/", http.StatusFound)
}

func (l *login) SignUp() {
	l.Ctx.Data["signdIn"] = signdIn(l.Ctx.Request(), l.Ctx.SessionStore)
	l.Ctx.Template = "login/signup"
	l.HTML(http.StatusOK)
}
