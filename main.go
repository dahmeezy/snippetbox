package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello from snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)
	w.Write([]byte(msg))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Create a new snippet..."))
}

func main() {
	Mux := http.NewServeMux()
	Mux.HandleFunc("GET /{$}", home)
	Mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	Mux.HandleFunc("GET /snippet/create", snippetCreate)

	log.Println("starting server on: 4000")
	err := http.ListenAndServe(": 4000", Mux)
	log.Println("This project was built by Zainab Yusuf")
	log.Fatal(err)
}
