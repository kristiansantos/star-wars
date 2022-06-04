package middlewares

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func Default(r *chi.Mux) {
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "x-api-key"},
		AllowCredentials: false,
	}

	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.StripSlashes)
	r.Use(cors.Handler(corsOptions))
	r.Use(middleware.Timeout(30 * time.Second))

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)

		response, _ := json.Marshal(map[string]interface{}{
			"error": map[string]interface{}{
				"status": http.StatusMethodNotAllowed,
				"title":  "method not allowed",
			},
		})

		w.Write([]byte(string(response)))
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		response, _ := json.Marshal(map[string]interface{}{
			"error": map[string]interface{}{
				"status": http.StatusNotFound,
				"title":  "not found",
			},
		})

		w.Write(response)
	})

	r.Get("/resource-status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response, _ := json.Marshal(map[string]interface{}{
			"status":  "ok",
			"data":    nil,
			"message": "Status ok",
		})

		w.Write(response)
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response, _ := json.Marshal(map[string]interface{}{
			"status":  "ok",
			"data":    nil,
			"message": "Status ok",
		})

		w.Write(response)
	})
}
