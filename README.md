# Go Middleware Demo

A simple Go HTTP server demonstrating middleware, handler, and resolver patterns with Docker support.

## Overview

This project showcases a clean architecture pattern commonly used in Go web applications:

- **Middleware**: Cross-cutting concerns like logging, authentication, etc.
- **Handler**: HTTP request/response handling and routing
- **Resolver**: Business logic and data processing

## Architecture

```
HTTP Request → Middleware → Handler → Resolver → Response
```

### Components

#### 1. Middleware (`middleware.go`)
- **LoggingMiddleware**: Logs incoming requests with timing information
- Wraps HTTP handlers to add functionality before/after request processing
- Demonstrates the middleware pattern for cross-cutting concerns

#### 2. Handler (`handler.go`)
- **GreetingHandler**: Processes HTTP requests and responses
- Extracts query parameters and coordinates with resolver
- Returns JSON responses

#### 3. Resolver (`resolver.go`)
- **GetGreeting**: Contains business logic for generating greetings
- Separates business logic from HTTP concerns
- Easy to test and reuse

## API Endpoints

### GET /greet
Returns a personalized greeting message.

**Query Parameters:**
- `name` (optional): Name to include in greeting. Defaults to "Guest"

**Example Requests:**
```bash
# Default greeting
curl http://localhost:8080/greet

# Personalized greeting
curl http://localhost:8080/greet?name=John
```

**Response:**
```json
{
  "message": "Hello, John! Welcome to Go Middleware Demo."
}
```

## Running the Application

### Local Development
```bash
go run .
```

### Using Docker

#### Build the image:
```bash
docker build -t go-middleware-demo .
```

#### Run the container:
```bash
docker run -p 8080:8080 go-middleware-demo
```

The server will start on port 8080 and you'll see logging output for each request.

## Project Structure

```
.
├── Dockerfile          # Multi-stage Docker build
├── README.md           # This file
├── go.mod              # Go module definition
├── main.go             # Application entry point and server setup
├── middleware.go       # Logging middleware implementation
├── handler.go          # HTTP request handlers
├── resolver.go         # Business logic layer
└── .gitignore          # Git ignore patterns
```

## Features

- **Clean Architecture**: Separation of concerns between middleware, handlers, and business logic
- **Logging Middleware**: Automatic request logging with timing
- **JSON API**: RESTful JSON responses
- **Docker Support**: Containerized deployment with multi-stage build
- **Lightweight**: Alpine-based Docker image (~14.7MB)

## Development

### Prerequisites
- Go 1.22 or later
- Docker (optional)

### Adding New Middleware
```go
func YourMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Before request processing
        
        next.ServeHTTP(w, r)
        
        // After request processing
    })
}
```

### Adding New Handlers
```go
func (h *Handler) YourHandler(w http.ResponseWriter, r *http.Request) {
    // Extract parameters
    // Call resolver
    // Return response
}
```

### Adding New Resolvers
```go
func (r *Resolver) YourBusinessLogic(params string) string {
    // Business logic here
    return result
}
```

## Example Output

When you make a request, you'll see logging output like:
```
2024/11/16 05:12:55 Server started on port 8080
2024/11/16 05:13:38 Incoming request → GET /greet
2024/11/16 05:13:38 Completed → GET /greet (616.083µs)
```

## License

This project is a demonstration example for educational purposes.