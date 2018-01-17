package controllers

import (
	"github.com/frzifus/dbwt/model"
	"github.com/gernest/utron/controller"
	"net/http"
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
