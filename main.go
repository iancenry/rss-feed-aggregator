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
