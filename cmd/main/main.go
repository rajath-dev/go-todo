package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rajath-dev/todo/pkg/routes"
)

const PORT = ":4000"

func main() {
	fmt.Println("Init")
	router := mux.NewRouter()
	routes.RegisterTodoRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
