package server

import "time"

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	Author    string    `json:"author"`
}

var posts = make(map[int]*Post)

var nextID int = 1

func createPost(title string, body string, author string) *Post {
	currentId := nextID
	p := Post{}
	p.ID = currentId
	p.Title = title
	p.Body = body
	p.CreatedAt = time.Now()
	p.Author = author
	nextID++

	posts[currentId] = &p

	return &p
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
