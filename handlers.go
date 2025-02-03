package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type UrlRequest struct {
	LongUrl     string `json:"long_url"`
	CustomAlias string `json:"custom_alias"`
	Expiration  int    `json:"expiration"`
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req UrlRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	shortID := req.CustomAlias
	if shortID == "" {
		shortID = generateShortID(5)
	}

	ctx := context.Background()

	expirationTime := time.Duration(req.Expiration) * time.Second
	err := RedisClient.Set(ctx, shortID, req.LongUrl, expirationTime).Err()
	if err != nil {
		http.Error(w, "Failed to store URL in Redis", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(map[string]string{"short_url": shortID})
	if err != nil {
		http.Error(w, "Error marshaling response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func generateShortID(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
