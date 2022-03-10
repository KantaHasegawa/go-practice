package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"

	"tutorial/controllers"
	"tutorial/models"
	"tutorial/utils"
)

func main() {
	router := mux.NewRouter()
	threadModel := models.Thread{}
	db := utils.ConnectDatabase("mysql", "root:@(localhost:3306)/go_practice")
	threadController := controller.ThreadController{Model: threadModel, Db: db}
	router.HandleFunc("/threads", threadController.ThreadIndexHandler).Methods("GET")
	router.HandleFunc("/thread/{id}", threadController.ThreadShowHandler).Methods("GET")
	router.HandleFunc("/thread/new", threadController.ThreadNewHandler).Methods("POST")
	http.Handle("/", router)
	fmt.Println("Server Start Listen to 8080")
	http.ListenAndServe(":8080", nil)
}
