package controller

import (
	"database/sql"
	"fmt"
	"net/http"

	"tutorial/models"
)

type ThreadController struct {
	Model models.Thread
	Db *sql.DB
}

func (threadController ThreadController) ThreadIndexHandler(w http.ResponseWriter, r *http.Request){
	result, err := threadController.Model.GetAll()
	if(err != nil){
		fmt.Println(err.Error())
		fmt.Fprintln(w, "Error: Sorry Server Error...")
		return
	}
	fmt.Fprintf(w, "%s\n", result)
}

func (threadController ThreadController) ThreadShowHandler(w http.ResponseWriter, r *http.Request){
	threadController.Model.Get()
	fmt.Println("show")
}

func (threadController ThreadController) ThreadNewHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("new")
}
