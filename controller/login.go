package controllers

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/frzifus/dbwt/model"
	u "github.com/frzifus/dbwt/util"
	"github.com/gernest/utron/controller"
)

type login struct {
	controller.BaseController

	Routes []string

	service *model.Service
}

// NewLogin -
func NewLogin() controller.Controller {
	return &login{
		Routes: []string{
			"get;/SignUp;SignUp",
			"get;/SignIn;SignIn",
			"get;/SignOff;SignOff",
			"get;/MyAccount;MyAccount",

			"post;/SignIn;SignIn",
			"post;/Register;Register",
			"post;/MyAccount;MyAccount",
		},
	}
}

func (l *login) encryptPassword(password string) string {
	h := sha256.Sum256([]byte(password))
	return "{SHA256}" + base64.StdEncoding.EncodeToString(h[:])
}

func (l *login) SignIn() {
	r := l.Ctx.Request()
	ln := r.FormValue("login_name")
	pwd := r.FormValue("password")
	pwd64 := l.encryptPassword(pwd)
	user := &model.User{}
	if err := l.Ctx.DB.Where("loginname = ? AND hash = ?",
		ln, pwd64).Find(&user).Error; err != nil {
		l.Ctx.Data["login_name"] = ln
		l.Ctx.Template = "login/signin"
		l.HTML(http.StatusOK)
		return
	}
	l.Ctx.DB.Model(&user).Update("LastLogin", time.Now())
	u.NewSession(l.Ctx, user)

	if referer := r.Header.Get("Referer"); !strings.Contains(referer, "SignIn") {
		l.Ctx.Redirect(referer, http.StatusFound)
	}
	l.Ctx.Data["signedIn"] = true
	l.Ctx.Template = "login/signin"
	l.HTML(http.StatusOK)
}

func (l *login) MyAccount() {
	l.service = model.NewService(l.Ctx)
	r := l.Ctx.Request()

	id, err := u.ReadID(l.Ctx)
	if err != nil {
		l.Ctx.Redirect(r.Header.Get("Referer"), http.StatusFound)
	}
	l.Ctx.Data["user"], _ = l.service.GetUserByID(id)
	l.Ctx.Data["role"] = u.Role(l.Ctx)
	l.Ctx.Template = "login/success"
	l.HTML(http.StatusOK)
}

func (l *login) SignOff() {
	u.DeleteSession(l.Ctx)
}

func (l *login) Register() {
	r := l.Ctx.Request()
	newUser := model.User{
		Active:    true,
		Firstname: r.FormValue("first_name"),
		Lastname:  r.FormValue("last_name"),
		Mail:      r.FormValue("email"),
		Loginname: r.FormValue("display_name"),
		Algo:      "sha256",
		Hash:      l.encryptPassword(r.FormValue("password")),
	}

	switch role := r.FormValue("role"); role {
	case "guest":
		if err := l.service.CreateGuest(newUser); err != nil {
			l.Ctx.DB.Rollback()
			fmt.Printf("Could not create user: %s", err)
			l.Ctx.Redirect("/SignUp?status=error", http.StatusFound)
		}
	case "student":
		if err := l.service.CreateStudent(newUser); err != nil {
			l.Ctx.DB.Rollback()
			fmt.Printf("Could not create user: %s", err)
			l.Ctx.Redirect("/SignUp?status=error", http.StatusFound)
		}
	case "employee":
		if err := l.service.CreateEmployee(newUser); err != nil {
			l.Ctx.DB.Rollback()
			fmt.Printf("Could not create user: %s", err)
			l.Ctx.Redirect("/SignUp?status=error", http.StatusFound)
		}
	}
	l.Ctx.Redirect("/SignUp?status=success", http.StatusFound)
}

func (l *login) SignUp() {
	r := l.Ctx.Request()
	if r.URL.Query().Get("status") == "success" {
		l.Ctx.Data["status"] = "success"
	} else if r.URL.Query().Get("status") == "error" {
		l.Ctx.Data["status"] = "error"
	} else {
		l.Ctx.Data["status"] = ""
	}
	l.Ctx.Template = "login/signup"
	l.HTML(http.StatusOK)
}
