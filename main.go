package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var prefixPath = "/api/compra"

// FindShoppingByIDController - Encuentra una compra por su ID
func FindShoppingByIDController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

// FindShoppingByUserController - Encuentra una compra por su ID
func FindShoppingByUserController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

// CreateShoppingController - Crear una compra
func CreateShoppingController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

// UpdateShoppingController - Actualiza una compra
func UpdateShoppingController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

// DeleteShoppingController - Borrr una compra
func DeleteShoppingController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc(prefixPath, CreateShoppingController).Methods("POST")
	r.HandleFunc(prefixPath, UpdateShoppingController).Methods("PUT")
	r.HandleFunc(prefixPath, DeleteShoppingController).Methods("DELETE")
	r.HandleFunc(prefixPath+"/{id}", FindShoppingByIDController).Methods("GET")
	r.HandleFunc(prefixPath+"/by-user/{id_user}", FindShoppingByUserController).Methods("GET")

	log.Printf("Listening...")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
