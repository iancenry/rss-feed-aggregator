package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load()
	port := os.Getenv("PORT") 

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
	// v1Router.HandleFunc("/healthz", readinessHandler)
	v1Router.Get("/healthz", readinessHandler)
	v1Router.Get("err", errorHandler)

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
