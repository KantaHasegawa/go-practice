package post

import (
	"fmt"
	"net/http"
)

func PostIndexHandler(http.ResponseWriter, *http.Request){
	fmt.Println("index")
}

func PostShowHandler(http.ResponseWriter, *http.Request){
	fmt.Println("show")
}
