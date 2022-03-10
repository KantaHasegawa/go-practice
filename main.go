package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"

	"tutorial/controllers"
)

// func router (path string)  {
// 	switch path {
// 	case "/":

// 	}
// }

// func handler (w http.ResponseWriter, r *http.Request) {
// 	fmt.Println(r.URL.Path)
// 	type Person struct {
// 		Name string `json:"name"`
// 		Age int `json:"age"`
// 	}
// 	person := Person{"kanta", 5}
// 	var buf bytes.Buffer
// 	enc := json.NewEncoder(&buf)
// 	if err := enc.Encode(person); err != nil {log.Fatal(err)}
// 	fmt.Fprint(w, buf.String())
// }

func main() {
	router := mux.NewRouter()
	threadController := new(controller.ThreadController)
	router.HandleFunc("/threads", threadController.ThreadIndexHandler).Methods("GET")
	router.HandleFunc("/thread/{id}", threadController.ThreadShowHandler).Methods("GET")
	router.HandleFunc("/thread/new", threadController.ThreadNewHandler).Methods("POST")
	http.Handle("/", router)
	fmt.Println("Server Start Listen to 8080")
	http.ListenAndServe(":8080", nil)
}
