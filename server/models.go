package server

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	Author    string    `json:"author"`
}

var posts = make(map[int]*Post)

var nextID int = 1

func createPost(title, body, author string) *Post {
	post := &Post{
		ID:        nextID,
		Title:     title,
		Body:      body,
		Author:    author,
		CreatedAt: time.Now(),
	}

	posts[nextID] = post
	nextID++

	err := savePosts()
	if err != nil {
		log.Printf("Failed to save posts: %v", err)
	}

	return post
}

func readAllPosts() []*Post {
	var s []*Post

	for _, value := range posts {
		s = append(s, value)

	}
	return s
}

func readPostById(id int) *Post {
	val, ok := posts[id]
	if !ok {
		return nil
	}

	return val

}

func updatePost(id int, body, title string) *Post {
	val, ok := posts[id]
	if !ok {
		return nil
	}
	val.Title = title
	val.Body = body

	err := savePosts()
	if err != nil {
		log.Printf("Failed to save posts: %v", err)
	}

	return val

}

func deletePost(id int) *Post {
	val, ok := posts[id]
	if !ok {
		return nil
	}

	delete(posts, id)

	err := savePosts()
	if err != nil {
		log.Printf("Failed to save posts: %v", err)
	}
	return val
}

func savePosts() error {
	data, err := json.Marshal(posts)
	if err != nil {
		return err
	}

	err = os.WriteFile("posts.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func LoadPosts() {
	data, err := os.ReadFile("posts.json")
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("No posts loaded yet")
			return
		}

		log.Printf("Failed to load posts: %v", err)
		return
	}

	err = json.Unmarshal(data, &posts)
	if err != nil {
		log.Printf("Failed to parse posts file: %v", err)
		return
	}

	maxID := 0
	for _, post := range posts {
		if post.ID > maxID {
			maxID = post.ID
		}
	}
	nextID = maxID + 1
}
