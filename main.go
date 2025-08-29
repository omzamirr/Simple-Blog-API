package main

import (
	"log"
	"net/http"

	"github.com/omzamirr/Simple-Blog-API/server"
)

func main() {
	http.HandleFunc("/posts", server.HandlePosts)  // For exact /posts
	http.HandleFunc("/posts/", server.HandlePosts) // For /posts/1, /posts/2, etc.
	log.Fatal(http.ListenAndServe(":8080", nil))
}
