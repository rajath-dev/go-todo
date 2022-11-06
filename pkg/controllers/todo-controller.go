package controllers

import (
	"fmt"
	"net/http"
)

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ALL TODOS")
}
