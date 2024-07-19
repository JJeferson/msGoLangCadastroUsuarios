package main

import (
	"log"
	"net/http"
	"project-root/rest/post"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/produzir", post.Produzir).Methods("POST")

	log.Println("Server started at :3030")
	log.Fatal(http.ListenAndServe(":3030", router))
}
