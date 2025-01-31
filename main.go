package main

import (
	"fmt"
	"net/http"
)

func main() {
	serveMux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":8080",
		Handler: serveMux,
	}

	handler := http.FileServer(http.Dir("."))
	serveMux.Handle("/", handler)

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("failed to run server.")
	}
}
