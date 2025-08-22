package routes

import (
	"net/http"
	"os"
	"strings"
	"github.com/stjudewashere/seonaut/internal/services"
)

// CORSMiddleware handles Cross-Origin Resource Sharing (CORS) for API requests
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get allowed origins from environment or use defaults
		allowedOrigins := getAllowedOrigins()
		
		origin := r.Header.Get("Origin")
		
		// Check if the origin is allowed
		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				break
			}
		}
		
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "86400") // 24 hours
		
		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}

// CORSHandler is a wrapper function for individual handlers
func CORSHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get allowed origins from environment or use defaults
		allowedOrigins := getAllowedOrigins()
		
		origin := r.Header.Get("Origin")
		
		// Check if the origin is allowed
		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				break
			}
		}
		
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "86400") // 24 hours
		
		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		handler.ServeHTTP(w, r)
	}
}

// getAllowedOrigins returns the list of allowed origins for CORS
func getAllowedOrigins() []string {
	// Default allowed origins
	defaultOrigins := []string{
		"https://codified-seo.vercel.app",
		"http://localhost:3000",
		"http://localhost:3001",
		"http://localhost:5173", // Vite dev server
		"http://localhost:8080", // Common dev server port
	}
	
	// Check for environment variable override
	if envOrigins := os.Getenv("SEONAUT_ALLOWED_ORIGINS"); envOrigins != "" {
		return strings.Split(envOrigins, ",")
	}
	
	// Check for additional frontend URL from environment
	if frontendURL := os.Getenv("SEONAUT_FRONTEND_URL"); frontendURL != "" {
		defaultOrigins = append(defaultOrigins, frontendURL)
	}
	
	return defaultOrigins
}
