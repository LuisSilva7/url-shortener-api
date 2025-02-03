package main

import (
	"net/http"
	"strings"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/shorten", ShortenHandler)

	mux.HandleFunc("/", DynamicHandler)

	return mux
}

func DynamicHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")

	if strings.HasPrefix(path, "stats/") {
		// StatsHandler(w, r)
	} else {
		// RedirectHandler(w, r)
	}
}

