package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  blog create <title> <content> <author>")
		fmt.Println("  blog list")
		fmt.Println("  blog show <id>")
		fmt.Println("  blog delete <id>")
		return
	}

	command := os.Args[1]

	switch command {
	case "create":
		if len(os.Args) < 5 {
			fmt.Println("Please provide title, content, and author")
			return
		}
		title := os.Args[2]
		content := os.Args[3]
		author := os.Args[4]

		jsonStr := fmt.Sprintf(`{"title":"%s","body":"%s","author":"%s"}`, title, content, author)
		resp, err := http.Post("http://localhost:8080/posts", "application/json", strings.NewReader(jsonStr))
		if err != nil {
			fmt.Printf("Error making request: %v\n", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response: %v\n", err)
			return
		}

		fmt.Println("Post created:")
		fmt.Println(string(body))

	case "list":
		resp, err := http.Get("http://localhost:8080/posts")
		if err != nil {
			fmt.Printf("Error making request: %v\n", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response: %v\n", err)
			return
		}

		fmt.Println("All posts:")
		fmt.Println(string(body))

	case "show":
		if len(os.Args) < 3 {
			fmt.Println("Please provide post ID")
			return
		}
		id := os.Args[2]

		resp, err := http.Get("http://localhost:8080/posts/" + id)
		if err != nil {
			fmt.Printf("Error making request: %v\n", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response: %v\n", err)
			return
		}

		fmt.Println("Post:")
		fmt.Println(string(body))

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide post ID")
			return
		}
		id := os.Args[2]

		req, err := http.NewRequest("DELETE", "http://localhost:8080/posts/"+id, nil)
		if err != nil {
			fmt.Printf("Error creating request: %v\n", err)
			return
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error making request: %v\n", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == 204 {
			fmt.Println("Post deleted successfully")
		} else {
			body, _ := io.ReadAll(resp.Body)
			fmt.Println(string(body))
		}

	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Use 'blog' without arguments to see usage")
	}
}
