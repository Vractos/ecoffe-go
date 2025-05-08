package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Vractos/ecoffe-go/adapter/api/handler"
	mdw "github.com/Vractos/ecoffe-go/adapter/api/middleware"
	"github.com/Vractos/ecoffe-go/adapter/repository"
	"github.com/Vractos/ecoffe-go/pkg/metrics"
	"github.com/Vractos/ecoffe-go/usecases/order"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func startEnv() {
	if env := os.Getenv("APP_ENV"); env == "" || env == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}
}

func main() {
	// ENV
	startEnv()

	// Log
	logger := metrics.NewLogger("info")
	defer logger.Sync()

	// PostgreSQL
	dataSourceName := fmt.Sprintf("postgresql://%s:%s@%s:5432/%s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_DB_NAME"))
	dbpool, err := pgxpool.New(context.Background(), dataSourceName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	// Repositories
	orderRepo := repository.NewOrderPostgreSQL(dbpool, *logger)

	// Services
	orderService := order.NewOrderService(
		orderRepo,
		logger,
	)

	// Router
	r := chi.NewRouter()
	r.Use(mdw.NewStructuredLogger(logger))
	// r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Public Routes
	r.Group(func(r chi.Router) {
		// "/order"
		handler.MakeOrderHandlers(r, orderService, *logger)
	})

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	logger.Info("Listing on 80")
	PORT := ":80"
	if env := os.Getenv("APP_ENV"); env == "" || env == "development" {
		PORT = ":8080"
	}
	err = http.ListenAndServe(PORT, r)
	if err != nil {
		logger.Panic(err.Error(), err)
	}
}
