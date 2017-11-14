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
	err = app.Model.Register(&m.Product{})
	if err != nil {
		log.Fatal(err)
	}
	err = app.Model.Register(&m.Test{})
	if err != nil {
		log.Fatal(err)
	}

	// CReate Models tables if they dont exist yet
	// app.Model.AutoMigrateAll()

	// Register Controllers
	app.AddController(c.NewIndex)
	app.AddController(c.NewProducts)

	// Start the server
	port := fmt.Sprintf(":%d", app.Config.Port)
	app.Log.Info("staring server on port", port)
	log.Fatal(http.ListenAndServe(port, app))
}
