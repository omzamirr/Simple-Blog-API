package server

import (
	"encoding/json"
	"net/http"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	read := readallPosts()
	data, err := json.Marshal(read)
	if err != nil {
		http.Error(w, "message", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
