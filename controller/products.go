package controllers

import (
	"github.com/frzifus/dbwt/model"
	"github.com/gernest/utron/controller"
	"net/http"
	"strconv"
)

type products struct {
	controller.BaseController
	Routes []string
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
	r := p.Ctx.Request()
	p.Ctx.Data["signedIn"] = signedIn(r, p.Ctx.SessionStore)
	p.Ctx.Data["title"] = "Produkte"
	p.Ctx.Data["products"], _ = p.getAllProducts()

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
	r := p.Ctx.Request()
	p.Ctx.Data["signedIn"] = signedIn(r, p.Ctx.SessionStore)
	productID := r.URL.Query().Get("product_id")

	id, err := strconv.Atoi(productID)

	product, perr := p.getProductByID(id)

	if err != nil || perr != nil {
		p.Ctx.Redirect("/Products?missing=1", http.StatusFound)
	}

	p.Ctx.Data["title"] = "Detail"
	p.Ctx.Data["product"] = product
	p.Ctx.Data["role"] = role(r, p.Ctx.SessionStore)
	p.Ctx.Template = "products/detail"
	p.HTML(http.StatusOK)
}

func (p *products) Ingredients() {
	p.Ctx.Data["signedIn"] = signedIn(p.Ctx.Request(), p.Ctx.SessionStore)
	p.Ctx.Data["title"] = "Ingredients"
	p.Ctx.Data["ingredients"], _ = p.getAllIngredients()
	p.Ctx.Template = "products/ingredients"
	p.HTML(http.StatusOK)
}

// DB - methods

func (p *products) getProductByID(id int) (*model.Product, error) {
	product := &model.Product{}
	if err := p.Ctx.DB.Preload("Price").Preload("Orders").Preload("Ingredients").First(&product, id).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (p *products) getAllProducts() ([]*model.Product, error) {
	products := []*model.Product{}
	if err := p.Ctx.DB.Preload("Price").Find(&products).Error; err != nil {
		return products, err
	}
	return products, nil
}

func (p *products) getAllIngredients() ([]*model.Ingredient, error) {
	ingredients := []*model.Ingredient{}
	if err := p.Ctx.DB.Find(&ingredients).Error; err != nil {
		return ingredients, err
	}
	return ingredients, nil
}
