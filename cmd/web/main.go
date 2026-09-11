package main

import (
	"log"
	"net/http"
)

func main() {
	Mux := http.NewServeMux()
	Mux.HandleFunc("GET /{$}", home)
	Mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	Mux.HandleFunc("GET /snippet/create", snippetCreate)
	Mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Println("starting server on: 4000")
	err := http.ListenAndServe("4000", Mux)
	log.Println("This project was built by Zainab Yusuf")
	log.Fatal(err)
}
