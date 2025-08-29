package server

import "time"

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
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

	return val

}

func deletePost(id int) *Post {
	val, ok := posts[id]
	if !ok {
		return nil
	}

	delete(posts, id)
	return val
}
