package main

import (
	"log"
	"net/http"
)

func main() {

	resolver := &Resolver{}
	handler := &Handler{Resolver: resolver}

	mux := http.NewServeMux()
	mux.Handle("/greet", LoggingMiddleware(http.HandlerFunc(handler.GreetingHandler)))

	log.Println("Server started on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
