# URL Shortener API

A simple URL shortener API built with Go and Redis, allowing users to shorten URLs, redirect to the original link, and track visit counts.

## Features

- **Shorten URLs**: Generate short links for long URLs.
- **Custom Aliases**: Users can specify a custom short link.
- **Redirection**: Automatically redirects to the original URL when accessing a short link.
- **Visit Tracking**: Tracks the number of visits to each short link.
- **Expiration**: Allows setting an expiration time for shortened URLs.

## Technologies Used

- **Go** (`net/http`, `encoding/json`, `github.com/redis/go-redis/v9`)
- **Redis** (used for storing short URLs and tracking visits)
- **Docker (optional)** (for running Redis in a container)

## How to Run the Project Locally

### Installation Steps

1. **Clone the repository:**
   ```bash
   git clone https://github.com/LuisSilva7/url-shortener-api.git
   ```

1. **Run redis on docker:**
   ```bash
   docker compose up -d
   ```
   
3. **Navigate to the project directory:**

   ```bash
   cd url-shortener-api
   ```

4. **Compile the project:**

   ```bash
   make build
   ```

5. **Run the server:**

   ```bash
   make run
   ```

The server will start and listen on port 8080. You can test it with:

  ```bash
  # Shorten a URL
  curl -X POST http://localhost:8080/shorten \
       -H "Content-Type: application/json" \
       -d '{"long_url": "https://example.com"}'
  
  # Shorten a URL with a custom alias
  curl -X POST http://localhost:8080/shorten \
       -H "Content-Type: application/json" \
       -d '{"long_url": "https://example.com", "custom_alias": "mylink"}'
  
  # Redirect using a short URL
  curl -X GET http://localhost:8080/mylink
  
  # Get visit statistics for a short URL
  curl -X GET http://localhost:8080/stats/mylink
  ```
