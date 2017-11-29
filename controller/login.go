package controllers

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/frzifus/dbwt/model"
	"github.com/gernest/utron/controller"
	"github.com/gorilla/schema"
	"net/http"
)

var decoder = schema.NewDecoder()

type login struct {
	controller.BaseController
	Routes []string
}

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
	l.Ctx.Data["signdIn"] = l.signedIn()
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
	l.Ctx.Data["user"] = u
	l.Ctx.Template = "login/success"
	l.HTML(http.StatusOK)
}

func (l *login) SignOff() {
	session, err := l.Ctx.SessionStore.Get(l.Ctx.Request(), "text")
	if err != nil {

	}
	fmt.Println(session.Name())

	l.Ctx.Redirect("/", http.StatusFound)
}

func (l *login) signedIn() bool {
	_, err := l.Ctx.SessionStore.New(l.Ctx.Request(), "test")
	if err != nil {
		fmt.Println(err)
	}
	return false
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

	// if err := decoder.Decode(newUser, req.PostForm); err != nil {
	//	l.Ctx.Data["Message"] = err.Error()
	//	l.Ctx.Template = "error"
	//	l.HTML(http.StatusInternalServerError)
	//	return
	// }
	if err := l.Ctx.DB.Create(newUser).Error; err != nil {
		l.Ctx.DB.Rollback()
		fmt.Printf("Could not create user: %s", err)
	}

	l.Ctx.Redirect("/", http.StatusFound)
}

func (l *login) SignUp() {
	l.Ctx.Data["SigndIn"] = l.signedIn()
	l.Ctx.Template = "login/signup"
	l.HTML(http.StatusOK)
}
