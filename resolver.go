package main

// Resolver simulates fetching data or doing business logic.
type Resolver struct{}

func (r *Resolver) GetGreeting(name string) string {
	if name == "" {
		name = "Guest"
	}
	return "Hello, " + name + "! Welcome to Go Middleware Demo."
}
