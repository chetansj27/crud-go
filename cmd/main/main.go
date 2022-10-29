package main

import (
	"log"
	"net/http"

	"github.com/chetansj27/crud-go/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()
	routes.RegisterUserRoutes(route)
	http.Handle("/", route)
	log.Fatal(http.ListenAndServe("localhost:8182", route))
}
