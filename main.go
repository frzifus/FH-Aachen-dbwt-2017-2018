package main

import (
	"fmt"
	"log"
	"net/http"

	c "git.klimlive.de/frzifus/dbwt/controller"
	m "git.klimlive.de/frzifus/dbwt/model"
	u "git.klimlive.de/frzifus/dbwt/util"
	"github.com/gernest/utron"
)

var version string

func main() {
	// Start the MVC Apo
	app, err := utron.NewMVC()
	if err != nil {
		log.Fatal(err)
	}

	err = app.SetNotFoundHandler(http.HandlerFunc(u.NotFoundHandler.ServeHttp))
	if err != nil {
		log.Fatal(err)
	}

	// Register Models
	err = app.Model.Register(
		&m.Product{},
		&m.User{},
		&m.Member{},
		&m.Guest{},
		&m.Student{},
		&m.Employee{},
		&m.Order{},
		&m.Image{},
		&m.Category{},
		&m.Price{},
		&m.Ingredient{},
		&m.Admin{},
	)
	if err != nil {
		log.Fatal(err)
	}

	// CReate Models tables if they dont exist yet
	fmt.Printf("Database: %s\n DatabaseConn: %s\n",
		app.Config.Database, app.Config.DatabaseConn)
	if err := app.Model.OpenWithConfig(app.Config); err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", app.Model.IsOpen())
	app.Model.SingularTable(false)
	app.Model.AutoMigrateAll()

	// Register Controllers
	app.AddController(c.NewIndex, u.IsSignedIn(), u.CtxTitle("Startseite"))
	app.AddController(c.NewProducts, u.IsSignedIn(), u.CtxTitle("Produkte"))
	app.AddController(c.NewLogin, u.IsSignedIn(), u.CtxTitle("Auth"))
	app.AddController(c.NewAdmin, u.IsSignedIn(), u.CtxTitle("Admin"))

	// Start the server
	app.Log.Info("running version ", version)
	port := fmt.Sprintf(":%d", app.Config.Port)
	app.Log.Info("staring server on port", port)
	log.Fatal(http.ListenAndServe(port, app))
}
