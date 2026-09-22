package main

import ("log"
	"net/http"
	"flag"
)

func main() {

	addr:=flag.String("addr",":4000","HTTP network address")
	flag.Parse()


	Mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))

	Mux.Handle("GET /static/",http.StripPrefix("/static",fileServer))
	Mux.HandleFunc("GET /{$}", home)
	Mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	Mux.HandleFunc("GET /snippet/create", snippetCreate)
	Mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Printf("starting server on %s",*addr)
	err := http.ListenAndServe(*addr, Mux)
	log.Println("This project was built by Zainab Yusuf")
	log.Fatal(err)
}
