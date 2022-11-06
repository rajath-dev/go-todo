package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rajath-dev/todo/pkg/config"
	"github.com/rajath-dev/todo/pkg/routes"
)

const PORT = ":4000"

func connectToDB() {
	config.ConnectDB()
}

func initRoutes() {
	router := mux.NewRouter()
	routes.RegisterTodoRoutes(router)
	http.Handle("/", router)
}

func handleError() {
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func main() {
	initRoutes()
	connectToDB()
	handleError()
}
