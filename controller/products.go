package controllers

import (
	"net/http"
	"strconv"

	"git.klimlive.de/frzifus/dbwt/model"
	u "git.klimlive.de/frzifus/dbwt/util"
	"github.com/gernest/utron/controller"
)

type products struct {
	controller.BaseController

	Routes []string

	service *model.Service
}

func NewProducts() controller.Controller {
	return &products{
		Routes: []string{
			"get;/Products;Products",
			"get;/Detail;Detail",
			"get;/Ingredients;Ingredients",
		},
	}
}

func (p *products) Products() {
	p.service = model.NewService(p.Ctx)
	r := p.Ctx.Request()
	p.Ctx.Data["products"], _ = p.service.GetAllProducts()

	missing, err := strconv.Atoi(r.URL.Query().Get("missing"))

	if missing > 0 && err == nil {
		p.Ctx.Data["missing"] = true
	} else {
		p.Ctx.Data["missing"] = false
	}

	p.Ctx.Template = "products/products"
	p.HTML(http.StatusOK)
}

func (p *products) Detail() {
	p.service = model.NewService(p.Ctx)
	r := p.Ctx.Request()
	productID := r.URL.Query().Get("product_id")

	id, err := strconv.Atoi(productID)

	product, perr := p.service.GetProductByID(id)

	if err != nil || perr != nil {
		p.Ctx.Redirect("/Products?missing=1", http.StatusFound)
	}

	p.Ctx.Data["title"] = "Detail"
	p.Ctx.Data["product"] = product
	p.Ctx.Data["role"] = u.Role(p.Ctx)
	p.Ctx.Template = "products/detail"
	p.HTML(http.StatusOK)
}

func (p *products) Ingredients() {
	p.service = model.NewService(p.Ctx)
	p.Ctx.Data["title"] = "Ingredients"
	p.Ctx.Data["ingredients"], _ = p.service.GetAllIngredients()
	p.Ctx.Template = "products/ingredients"
	p.HTML(http.StatusOK)
}
