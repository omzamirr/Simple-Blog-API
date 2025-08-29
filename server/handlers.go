package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	read := readAllPosts()
	data, err := json.Marshal(read)
	if err != nil {
		http.Error(w, "message", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Could not read request body", http.StatusBadRequest)
		return
	}

	log.Printf("Received JSON: %s", string(data))

	var post Post
	err = json.Unmarshal(data, &post)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	createdPost := createPost(post.Title, post.Body, post.Author)
	data, err = json.Marshal(createdPost)
	if err != nil {
		http.Error(w, "Could not create response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

func HandlePosts(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		GetPosts(w, r)
	} else if r.Method == "POST" {
		CreatePost(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
