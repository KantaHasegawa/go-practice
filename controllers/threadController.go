package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"tutorial/models"
	"tutorial/utils"

	"github.com/gorilla/mux"
)

type ThreadController struct {
	Model models.ThreadInterface
	Db utils.DatabaseInterface
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
	body, err := io.ReadAll(r.Body)
	if(err != nil){
		fmt.Println(err.Error())
		fmt.Fprintln(w, "Error: Sorry Server Error...")
		return
	}
	thread := models.Thread{}
	err = json.Unmarshal(body, &thread)
	if(err != nil){
		fmt.Println(err.Error())
		fmt.Fprintln(w, "Error: Sorry Server Error...")
		return
	}
	fmt.Println(thread)
}
