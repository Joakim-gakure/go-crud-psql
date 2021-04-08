package main

import (
	"fmt"
	"go-crud-psql/pkg/db"
	handler "go-crud-psql/pkg/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	DB := db.InitializeDB()

	r.HandleFunc("/", handler.GetUser(DB)).Methods("GET")
	r.HandleFunc("/", handler.CreateUser(DB)).Methods("POST")

	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
