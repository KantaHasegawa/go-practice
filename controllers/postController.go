package controller

import (
	"fmt"
	"net/http"
)

type PostController struct {}

func (postController PostController) PostIndexHandler(http.ResponseWriter, *http.Request){
	fmt.Println("index")
}

func (postController PostController) PostShowHandler(http.ResponseWriter, *http.Request){
	fmt.Println("show")
}
