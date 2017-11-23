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
			"get;/Login;Login",
		},
	}
}

func (l *login) encryptPassword(password string) string {
	h := sha256.Sum256([]byte(password))
	return "{SHA256}" + base64.StdEncoding.EncodeToString(h[:])
}

// func (l *login) register() {
//	r := l.Ctx.Request()
//	newUser := model.Nutzer{
//		Aktiv:       true,
//		Vorname:     r.FormValue("vorname"),
//		Nachname:    r.FormValue("nachname"),
//		Mail:        r.FormValue("mail"),
//		Loginname:   r.FormValue("usr"),
//		Algorithmus: "sha256",
//		Hash:        l.encryptPassword(r.FormValue("pwd")),
//	}
//	if err := l.Ctx.DB.Create(&newUser).Error; err != nil {
//		l.Ctx.DB.Rollback()
//		fmt.Printf("Could not create user: %s", err)
//	}
// }

func (l *login) signIn() {
	r := l.Ctx.Request()
	usr := r.FormValue("usr")
	pwd := r.FormValue("pwd")
	pwd64 := l.encryptPassword(pwd)
	u := model.Nutzer{
		Loginname: usr,
		Hash:      pwd64,
	}
	l.Ctx.DB.Where(&u).First(&u)
	// if u.Aktiv {
	// session cookie setzen
	// session, _ := l.Ctx.NewSession("test")
	// l.Ctx.SaveSession(session)
	// }
}

func (l *login) signedIn() bool {
	_, err := l.Ctx.GetSession("test")
	if err != nil {
		fmt.Println(err)
	}
	return false
}

func (l *login) newSession(name string) error {
	session, _ := l.Ctx.NewSession(name)
	return l.Ctx.SaveSession(session)
}

func (l *login) Login() {
	l.signIn()

	if l.signedIn() {
		fmt.Println("success")
	}

	l.Ctx.Template = "login/login"
	l.HTML(http.StatusOK)
}
