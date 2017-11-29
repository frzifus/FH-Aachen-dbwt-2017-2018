package controllers

import (
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
    // check if logged in

	a.Ctx.Template = "admin/admin"
	a.HTML(http.StatusOK)
}

