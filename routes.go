package main

import (
	"net/http"
	"strings"
)

// SetupRoutes initializes the HTTP routes and returns a ServeMux
func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Route for shortening URLs
	mux.HandleFunc("/shorten", ShortenHandler)

	// Dynamic route handling for shortened URLs and stats
	mux.HandleFunc("/", DynamicHandler)

	return mux
}

// DynamicHandler handles requests that depend on URL parameters
func DynamicHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")

	// Check if the path matches a shortened URL or a stats request
	if strings.HasPrefix(path, "stats/") {
		StatsHandler(w, r)
	} else {
		RedirectHandler(w, r)
	}
}

