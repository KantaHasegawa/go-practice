package controller

import (
	"fmt"
	"net/http"

	"tutorial/models"
)

type ThreadController struct {
	model models.Thread
}

func (threadController ThreadController) ThreadIndexHandler(w http.ResponseWriter, r *http.Request){
	result, err := threadController.model.GetAll()
	if(err != nil){
		fmt.Println(err.Error())
		fmt.Fprintln(w, "Error: Sorry Server Error...")
		return
	}
	fmt.Fprintf(w, "%s\n", result)
}

func (threadController ThreadController) ThreadShowHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("show")
}

func (threadController ThreadController) ThreadNewHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("new")
}
