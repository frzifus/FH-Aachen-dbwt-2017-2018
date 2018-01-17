package controllers

import (
	"github.com/frzifus/dbwt/model"
	"github.com/gernest/utron/controller"
	"net/http"
	"time"
)

type admin struct {
	controller.BaseController
	Routes []string
}

func NewAdmin() controller.Controller {
	return &admin{
		Routes: []string{
			"get;/Admin;Admin",
		},
	}
}

func (a *admin) Admin() {
	if !a.isAdmin() {
		a.Ctx.Redirect("/", http.StatusFound)
	}
	a.Ctx.Data["signedIn"] = signedIn(a.Ctx.Request(), a.Ctx.SessionStore)

	layout := "2006-01-02 15:04:05"
	str := "0001-01-01 00:00:00"
	t, _ := time.Parse(layout, str)

	users := []*model.User{}
	if err := a.Ctx.DB.Where("last_login != ?", t).Find(&users).Error; err == nil {
		a.Ctx.Data["userLogins"] = users
	}
	if err := a.Ctx.DB.Where("last_login = ?", t).Find(&users).Error; err == nil {
		a.Ctx.Data["newRegistrations"] = users
	}

	a.Ctx.Template = "admin/admin"
	a.HTML(http.StatusOK)
}

func (a *admin) isAdmin() bool {
	admin := &model.Admin{}
	id, readErr := readID(a.Ctx.Request(), a.Ctx.SessionStore)
	if dbErr := a.Ctx.DB.Where("user_id = ?", id).First(&admin).Error; readErr != nil || dbErr != nil {
		return false
	}
	return true
}
