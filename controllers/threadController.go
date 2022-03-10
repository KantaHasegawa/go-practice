package controller

import (
	"fmt"
	"net/http"
)

type ThreadController struct {}

func (threadController ThreadController) ThreadIndexHandler(http.ResponseWriter, *http.Request){
	fmt.Println("index")
}

func (threadController ThreadController) ThreadShowHandler(http.ResponseWriter, *http.Request){
	fmt.Println("show")
}

func (threadController ThreadController) ThreadNewHandler(http.ResponseWriter, *http.Request){
	fmt.Println("new")
}
