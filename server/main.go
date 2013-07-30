package main

import (
	"./controllers"
	"./models"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stephenalexbrowne/zoom"
	"log"
	"net/http"
)

func main() {
	port := "6060"

	err := models.Initialize()
	if err != nil {
		log.Fatal(err)
	}
	defer zoom.Close()

	r := route()

	http.Handle("/", r)
	portName := ":" + port
	fmt.Printf("server listening on port %s...\n", port)
	http.ListenAndServe(portName, nil)
}

func route() *mux.Router {
	r := mux.NewRouter()

	// items
	itemsController := new(controllers.ItemsController)
	r.HandleFunc("/items", itemsController.Index).Methods("GET")
	r.HandleFunc("/items", itemsController.Create).Methods("POST")
	r.HandleFunc("/items/{id}", itemsController.Update).Methods("PUT")
	r.HandleFunc("/items/{id}", itemsController.Show).Methods("GET")
	r.HandleFunc("/items/{id}", itemsController.Delete).Methods("DELETE")

	return r
}
