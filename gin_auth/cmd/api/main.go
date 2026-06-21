package main

import (
	"context"
	"gin_auth/internal/app"
	"gin_auth/internal/httpserver"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background() // root context

	app, err := app.New(ctx)
	if err != nil {
		log.Fatalf("failed to start: %v", err)
	}
	defer func() {
		if err := app.Close(ctx); err != nil {
			log.Printf("Shutdown warning: %v", err)
		}
	}()

	router := httpserver.NewRouter(app)

	// builder server struct explictly
	server := &http.Server{
		Addr: ":3000",
		Handler: router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("API running on %s", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed{
			log.Printf("server closed")
			return
		}
		log.Fatalf("server error: %v", err)

	}
}