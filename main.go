package main

import (
	"log"
	"net/http"

	"github.com/omzamirr/Simple-Blog-API/server"
)

func main() {
	http.HandleFunc("/posts", server.HandlePosts)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
