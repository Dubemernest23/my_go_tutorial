package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.DefaultLogger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.NoCache)
	r.Use(middleware.AllowContentEncoding("deflate", "gzip"))
	r.Use(middleware.AllowContentType("application/json", "text/xml"))
	r.Use(middleware.CleanPath) // cleans paths mistakes of double slashes - user//u/1
	r.Use(middleware.Heartbeat("/"))
	r.Use(middleware.Throttle(15))
	r.Use(middleware.ThrottleBacklog(10, 50, time.Second*10))
	r.Use(middleware.Timeout(time.Second * 60))

	// Basic CORS
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// r.Mount("/debug", middleware.Profiler())
	// This line enables profiling
	r.Mount("/debug", middleware.Profiler())

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {})

	r.Get("/long", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		processTime := time.Duration(rand.Intn(4)+1) * time.Second

		select {
		case <-ctx.Done():
			return

		case <-time.After(processTime):
			// The above channel simulates some hard work.
		}

		w.Write([]byte("done"))
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("Not found"))
	})
	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		w.Write([]byte("method is not valid"))
	})
	http.ListenAndServe(":3000", r)
}
