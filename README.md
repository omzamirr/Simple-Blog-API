# CLIBlog - Personal Command-Line Blog System

A lightweight, file-based blogging system built in Go that provides both a REST API server and command-line interface for managing blog posts.

## Features

- **Complete CRUD Operations**: Create, read, update, and delete blog posts
- **REST API Server**: HTTP endpoints for programmatic access
- **Command-Line Interface**: Easy-to-use CLI for daily blogging
- **File-Based Persistence**: Posts stored in JSON format, no database required
- **Data Persistence**: Posts survive server restarts
- **Error Handling**: Comprehensive error handling throughout the system

## Project Structure

```
CLIBlog/
├── main.go                 # API server entry point
├── go.mod                  # Go module definition
├── posts.json              # Data storage (auto-generated)
├── server/
│   ├── handlers.go         # HTTP request handlers
│   └── models.go           # Data models and persistence
└── cmd/
    └── blog/
        └── main.go         # CLI client
```

## Quick Start

### 1. Start the API Server

```bash
go run main.go
```

The server will start on `http://localhost:8080` and load any existing posts from `posts.json`.

### 2. Use the CLI to Manage Posts

```bash
# Create a new post
go run cmd/blog/main.go create "My First Post" "This is the content of my post" "Author Name"

# List all posts
go run cmd/blog/main.go list

# Show a specific post by ID
go run cmd/blog/main.go show 1

# Delete a post
go run cmd/blog/main.go delete 1
```

## CLI Commands

### Create a Post
```bash
go run cmd/blog/main.go create <title> <content> <author>
```
**Example:**
```bash
go run cmd/blog/main.go create "Weekend Adventures" "Went hiking in the mountains today. The weather was perfect!" "John"
```

### List All Posts
```bash
go run cmd/blog/main.go list
```
Displays all posts in JSON format.

### Show Specific Post
```bash
go run cmd/blog/main.go show <post_id>
```
**Example:**
```bash
go run cmd/blog/main.go show 1
```

### Delete a Post
```bash
go run cmd/blog/main.go delete <post_id>
```
**Example:**
```bash
go run cmd/blog/main.go delete 1
```

## REST API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/posts` | Retrieve all posts |
| GET | `/posts/{id}` | Retrieve specific post by ID |
| POST | `/posts` | Create a new post |
| PUT | `/posts/{id}` | Update existing post |
| DELETE | `/posts/{id}` | Delete a post |

### API Examples

#### Create a Post
```bash
curl -X POST http://localhost:8080/posts \
  -H "Content-Type: application/json" \
  -d '{"title":"API Test","body":"Created via API","author":"Developer"}'
```

#### Get All Posts
```bash
curl http://localhost:8080/posts
```

#### Get Specific Post
```bash
curl http://localhost:8080/posts/1
```

#### Update a Post
```bash
curl -X PUT http://localhost:8080/posts/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Updated Title","body":"Updated content","author":"Developer"}'
```

#### Delete a Post
```bash
curl -X DELETE http://localhost:8080/posts/1
```

## Post Structure

Each post contains the following fields:

```json
{
  "id": 1,
  "title": "Post Title",
  "body": "Post content goes here",
  "author": "Author Name",
  "created_at": "2025-08-30T12:00:00Z"
}
```

- `id`: Unique identifier (auto-generated)
- `title`: Post title
- `body`: Post content
- `author`: Post author
- `created_at`: Creation timestamp (auto-generated)

## Data Storage

Posts are stored in `posts.json` in the project root directory. The file is automatically created when you create your first post. The JSON structure stores posts as a map with string keys (post IDs) and post objects as values.

## Installation

### Prerequisites
- Go 1.19 or later

### Setup
1. Clone the repository
2. Navigate to the project directory
3. Run the server: `go run main.go`
4. Use CLI commands from the project root directory

### Building Executables
```bash
# Build API server
go build -o blog-server main.go

# Build CLI client
go build -o blog cmd/blog/main.go
```

## Architecture

### API Server
- **handlers.go**: HTTP request handlers for each endpoint
- **models.go**: Data models, CRUD operations, and file persistence
- **main.go**: Server setup and routing

### CLI Client
- **cmd/blog/main.go**: Command-line interface that makes HTTP requests to the API server

### Data Flow
1. CLI commands make HTTP requests to the API server
2. API server processes requests and calls model functions
3. Model functions perform CRUD operations on in-memory data
4. Data changes are automatically saved to `posts.json`
5. On server startup, data is loaded from `posts.json` into memory

## Error Handling

The system includes comprehensive error handling:
- Invalid CLI arguments
- Network connection errors
- File I/O errors
- JSON parsing errors
- HTTP request/response errors
- Post not found errors

## Contributing

This is a personal project built for learning Go and REST API development. The codebase follows Go conventions and includes:
- Proper error handling
- Clean separation of concerns
- RESTful API design
- Command-line best practices

## License

This project is for educational purposes. Feel free to use and modify as needed.
