package main

import (
	"auth_project/internal/app"
	httpserver "auth_project/internal/httpServer"
	"context"
	"log"
	"net/http"
	"time"
)

func main() {

	// root context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	ctx := context.Background()
	a, err := app.New(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
		return
	}
	defer func() {
		if err := a.Close(ctx); err != nil {
			log.Printf("Error closing app: %v", err)
		}
	}()

	router := httpserver.NewRouter(a)

	srv := &http.Server{
		Addr:              ":5020",
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("Starting server on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server error: %v", err)
		return
	}
}
