package main

import "net/http"

func main() {

import "net/http"

	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	mux.Handle("/", http.FileServer(http.Dir(".")))
	mux.Handle("/assets", http.FileServer(http.Dir(".")))

	err := server.ListenAndServe()
	if err != nil {
		return
	}

}
