package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", api)
	server := &http.Server{
		Addr: ":8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func api(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, my first go api"))
}
