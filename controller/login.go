package controllers

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/frzifus/dbwt/model"
	"github.com/gernest/utron/controller"
	"net/http"
	"strings"
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
	if signedIn := signedIn(r, l.Ctx.SessionStore); signedIn {
		l.Ctx.Redirect(r.Referer(), http.StatusFound)
	}

	ln := r.FormValue("login_name")
	pwd := r.FormValue("password")
	pwd64 := l.encryptPassword(pwd)
	u := &model.User{}
	if err := l.Ctx.DB.Where("loginname = ? AND hash = ?",
		ln, pwd64).Find(&u).Error; err != nil {
		l.Ctx.Data["login_name"] = ln
		l.Ctx.Template = "login/signin"
		l.HTML(http.StatusOK)
		return
	}
	l.newSession(u)

	if referer := r.Header.Get("Referer"); !strings.Contains(referer, "SignIn") {
		l.Ctx.Redirect(referer, http.StatusFound)
	}
	l.Ctx.Data["signedIn"] = true
	l.Ctx.Template = "login/signin"
	l.HTML(http.StatusOK)
}

func (l *login) MyAccount() {
	r := l.Ctx.Request()
	if signedIn := signedIn(r, l.Ctx.SessionStore); !signedIn {
		l.Ctx.Redirect("/", http.StatusFound)
	}

	id, err := readID(r, l.Ctx.SessionStore)
	if err != nil {
		l.Ctx.Redirect(r.Header.Get("Referer"), http.StatusFound)
	}
	l.Ctx.Data["user"], _ = l.userByID(id)
	l.Ctx.Data["signedIn"] = signedIn(r, l.Ctx.SessionStore)
	l.Ctx.Data["role"] = role(r, l.Ctx.SessionStore)
	l.Ctx.Template = "login/success"
	l.HTML(http.StatusOK)
}

func (l *login) SignOff() {
	session, err := l.Ctx.SessionStore.Get(l.Ctx.Request(), cookieName)
	if err != nil {
		l.Ctx.Redirect("/", http.StatusFound)
	}
	session.Options.MaxAge = -1
	r := l.Ctx.Request()
	_ = session.Save(r, l.Ctx.Response())
	l.Ctx.Redirect(r.Header.Get("Referer"), http.StatusFound)
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
		if err := l.dbCreateGuest(newUser); err != nil {
			l.Ctx.DB.Rollback()
			fmt.Printf("Could not create user: %s", err)
			l.Ctx.Redirect("/SignUp?status=error", http.StatusFound)
		}
	case "student":
		if err := l.dbCreateStudent(newUser); err != nil {
			l.Ctx.DB.Rollback()
			fmt.Printf("Could not create user: %s", err)
			l.Ctx.Redirect("/SignUp?status=error", http.StatusFound)
		}
	case "employee":
		if err := l.dbCreateEmployee(newUser); err != nil {
			l.Ctx.DB.Rollback()
			fmt.Printf("Could not create user: %s", err)
			l.Ctx.Redirect("/SignUp?status=error", http.StatusFound)
		}
	}
	l.Ctx.Redirect("/SignUp?status=success", http.StatusFound)
}

func (l *login) SignUp() {
	l.Ctx.Data["signedIn"] = signedIn(l.Ctx.Request(), l.Ctx.SessionStore)
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

func (l *login) newSession(u *model.User) {
	session, _ := l.Ctx.NewSession(cookieName)
	session.Values["username"] = u.Loginname
	session.Values["id"] = u.ID
	session.Values["role"] = l.dbRole(u)
	session.Values["active"] = true
	session.Options.Path = "/"
	session.Options.MaxAge = 10 * 24 * 3600
	_ = session.Save(l.Ctx.Request(), l.Ctx.Response())
}

func (l *login) dbCreateUser(u *model.User) error {
	return l.Ctx.DB.Create(u).Error
}

func (l *login) dbCreateMember(m *model.Member) error {
	return l.Ctx.DB.Create(m).Error
}

func (l *login) dbCreateGuest(u model.User) error {
	guest := &model.Guest{}
	guest.User = u
	return l.Ctx.DB.Create(&guest).Error
}

func (l *login) dbCreateStudent(u model.User) error {
	if err := l.dbCreateUser(&u); err != nil {
		return err
	}
	m := model.Member{
		UserID: u.ID,
		User:   u,
	}
	if err := l.dbCreateMember(&m); err != nil {
		return err
	}
	s := model.Student{
		MemberID: m.UserID,
	}
	return l.Ctx.DB.Create(&s).Error
}

func (l *login) dbCreateEmployee(u model.User) error {
	if err := l.dbCreateUser(&u); err != nil {
		return err
	}
	m := model.Member{
		UserID: u.ID,
		User:   u,
	}
	if err := l.dbCreateMember(&m); err != nil {
		return err
	}
	e := model.Employee{
		MemberID: m.UserID,
	}
	return l.Ctx.DB.Create(&e).Error
}

func (l *login) dbRole(u *model.User) string {
	if _, err := l.guestByID(u.ID); err == nil {
		return "guest"
	} else if _, err := l.studentByUserID(u.ID); err == nil {
		return "student"
	} else if _, err := l.employeeByUserID(u.ID); err == nil {
		return "employee"
	}
	return ""
}

func (l *login) userByID(id uint) (*model.User, error) {
	u := &model.User{}
	err := l.Ctx.DB.First(&u, id).Error
	return u, err
}

func (l *login) studentByUserID(id uint) (*model.Student, error) {
	s := &model.Student{}
	err := l.Ctx.DB.Where("member_id = ?", id).First(&s).Error
	return s, err
}

func (l *login) guestByID(id uint) (*model.Guest, error) {
	g := &model.Guest{}
	err := l.Ctx.DB.First(&g, id).Error
	return g, err
}

func (l *login) employeeByUserID(id uint) (*model.Employee, error) {
	e := &model.Employee{}
	err := l.Ctx.DB.Where("member_id = ?", id).First(&e).Error
	return e, err
}
