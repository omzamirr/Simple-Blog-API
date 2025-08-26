package main

import (
	"github.com/omzamirr/Simple-Blog-API/server"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/posts", server.GetPosts)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
