package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	router := http.NewServeMux()
	fs := http.FileServer(http.Dir("./client/dist"))
	router.Handle("GET /", fs)

	port := ":8080"
	fmt.Printf("Server is listening on port %v", port)
	err := http.ListenAndServe(port, router)

	if err != nil {
		log.Fatal(err)
	}
}
