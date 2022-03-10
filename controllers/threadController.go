package controller

import (
	"fmt"
	"net/http"

	"tutorial/models"

	"github.com/go-gorp/gorp"
	"github.com/gorilla/mux"
)

type ThreadController struct {
	Model models.Thread
	Db *gorp.DbMap
}

func (threadController ThreadController) ThreadIndexHandler(w http.ResponseWriter, r *http.Request){
	result, err := threadController.Model.GetAll(threadController.Db)
	if(err != nil){
		fmt.Println(err.Error())
		fmt.Fprintln(w, "Error: Sorry Server Error...")
		return
	}
	fmt.Fprintf(w, "%s\n", result)
}

func (threadController ThreadController) ThreadShowHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	result, err := threadController.Model.Get(vars["id"], threadController.Db)
	if(err != nil){
		fmt.Println(err.Error())
		fmt.Fprintln(w, "Error: Sorry Server Error...")
		return
	}
	fmt.Fprintf(w, "%s\n", result)
}

func (threadController ThreadController) ThreadNewHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("new")
}
