package main

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	Resolver *Resolver
}

func (h *Handler) GreetingHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	message := h.Resolver.GetGreeting(name)

	resp := map[string]string{
		"message": message,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
