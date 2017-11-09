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
    // r := h.Ctx.Request()
    // user := r.URL.Query().Get("user")
    i.Ctx.Data["title"] = "dbwt"
    i.Ctx.Data["user"] = "bla"
    i.Ctx.Template = "index/home"
    i.HTML(http.StatusOK)
}
