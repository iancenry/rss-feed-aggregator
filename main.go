package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load()
	port := os.Getenv("PORT") 

	router := chi.NewRouter()

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
