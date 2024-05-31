package domain

import (
	"net/http"
	"strings"
)

func RootPathMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if path == "/" {
			http.ServeFile(w, r, "public/index.html")
			return
		}

		if strings.Contains(path, ".js") ||
			strings.Contains(path, ".css") ||
			strings.Contains(path, ".ico") {

			http.ServeFile(w, r, "public/"+path)
			return
		}

		if !strings.Contains(path, "/sale") && !strings.Contains(path, "/location") {
			http.ServeFile(w, r, "public/index.html")
			return
		}

		next.ServeHTTP(w, r)
	})
}

/**
 * Middleware to handle CORS
 * @param next http.Handler
 * @return http.Handler
 */
func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
