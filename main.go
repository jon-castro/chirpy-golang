package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	mux.Handle("/", http.FileServer(http.Dir(".")))

	err := server.ListenAndServe()
	if err != nil {
		return
	}

}
