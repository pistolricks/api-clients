package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/pistolricks/api-clients/internal/database"
	"github.com/pistolricks/api-clients/internal/handlers"
	"github.com/pistolricks/api-clients/internal/repository"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	// Initialize database
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.CloseDB()

	// Create tables
	if err := database.CreateTables(); err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}

	// Create repository
	schoolRepo := repository.NewSchoolRepository(database.DB)

	// Create handlers
	schoolHandler := handlers.NewSchoolHandler(schoolRepo)

	// Create router
	r := mux.NewRouter()

	// API routes
	api := r.PathPrefix("/api").Subrouter()

	// Schools routes
	schools := api.PathPrefix("/schools").Subrouter()
	schools.HandleFunc("", schoolHandler.GetSchools).Methods("GET")
	schools.HandleFunc("", schoolHandler.CreateSchool).Methods("POST")
	schools.HandleFunc("/{id:[0-9]+}", schoolHandler.GetSchool).Methods("GET")
	schools.HandleFunc("/{id:[0-9]+}", schoolHandler.UpdateSchool).Methods("PUT")
	schools.HandleFunc("/{id:[0-9]+}", schoolHandler.DeleteSchool).Methods("DELETE")
	schools.HandleFunc("/import", schoolHandler.ImportGeoJSON).Methods("POST")

	// Set up server
	port := getEnv("PORT", "8080")
	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Server starting on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Server shutting down...")

	// Create a deadline for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown the server
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
}

// Helper function to get environment variable with fallback
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
