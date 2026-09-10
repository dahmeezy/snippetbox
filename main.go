package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello from snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Create a new snippet..."))
}

func main() {
	Mux := http.NewServeMux()
	Mux.HandleFunc("/{$}", home)
	Mux.HandleFunc("/snippet/view/{id}", snippetView)
	Mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("starting server on: 4000")
	err := http.ListenAndServe(": 4000", Mux)
	log.Println("This project was built by Zainab Yusuf")
	log.Fatal(err)
}
