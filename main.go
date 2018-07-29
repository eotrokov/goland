package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	people "./controllers"
)

// The person Type (more like an object)

// Display all from the people var

// main function to boot up everything
func main() {
	inicialize()
	router := mux.NewRouter()
	router.HandleFunc("/people", people.GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", people.GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", people.CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", people.DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func inicialize() {
	people.LoadData()
}
