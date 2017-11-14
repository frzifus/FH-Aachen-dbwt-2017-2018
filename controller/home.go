package controllers

import (
	"net/http"

	"github.com/gernest/utron/controller"
)

type index struct {
	controller.BaseController
	Routes []string
}

func NewIndex() controller.Controller {
	return &index{
		Routes: []string{
			"get;/;Index",
			"get;/Home;Home",
		},
	}
}

func (i *index) Index() {
	i.Ctx.Redirect("/Home", http.StatusFound)
}

func (i *index) Home() {
	i.Ctx.Data["title"] = "Home"
	i.Ctx.Template = "index/home"
	i.HTML(http.StatusOK)
}
