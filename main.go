package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/iancenry/rss-feed-aggregator/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// holds a connection to a DB
type apiConfig struct {
	DB *database.Queries
} 

func main(){
	godotenv.Load()
	port := os.Getenv("PORT") 
	if port == "" {
		log.Fatal("PORT isn't found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL isn't found in the environment")
	}

	conn, er :=  sql.Open("postgres", dbURL)
	if er != nil {
		log.Fatal("Can't connect to database")
	}


	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", readinessHandler)
	v1Router.Get("/err", errorHandler)

	// nesting a v1 router under the /v1 path - full path /v1/healthz
	router.Mount("/v1", v1Router)

	server := &http.Server{
		Addr: ":"+port,
		Handler: router,
	}

	log.Printf("Listening on port %s", port)
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
