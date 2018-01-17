package main

import (
	"fmt"
	"log"
	"net/http"

	c "github.com/frzifus/dbwt/controller"
	m "github.com/frzifus/dbwt/model"
	"github.com/gernest/utron"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Oops! 404"))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Start the MVC Apo
	app, err := utron.NewMVC()
	if err != nil {
		log.Fatal(err)
	}

	err = app.SetNotFoundHandler(http.HandlerFunc(notFound))
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
	app.AddController(c.NewIndex)
	app.AddController(c.NewProducts)
	app.AddController(c.NewLogin)
	app.AddController(c.NewAdmin)

	// Start the server
	port := fmt.Sprintf(":%d", app.Config.Port)
	app.Log.Info("staring server on port", port)
	log.Fatal(http.ListenAndServe(port, app))
}
