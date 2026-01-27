package main

import (
	"fmt"
	"net/http"
)

func handlerURL(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Encurtador de URL Fisrt step!")
}

func main() {
	http.HandleFunc("/", handlerURL)
	println("Servidor rodando na porta http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
