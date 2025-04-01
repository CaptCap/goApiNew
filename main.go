package main

import "net/http"

func main() {
	http.HandleFunc("/", api)
}

func api(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, my first go api"))
}
