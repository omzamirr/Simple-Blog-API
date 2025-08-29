package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
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

func ReadPostById(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/posts/")
	strToInt, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	postID := readPostById(int(strToInt))

	if postID == nil {
		http.Error(w, "Post does not exist", http.StatusNotFound)
		return
	}

	data, err := json.Marshal(postID)
	if err != nil {
		http.Error(w, "Could not create response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}

func HandlePosts(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if r.URL.Path == "/posts" {
			GetPosts(w, r)
		} else {
			ReadPostById(w, r)
		}
	} else if r.Method == "POST" {
		CreatePost(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
