package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load()
	port := os.Getenv("PORT")

	if err != nil{
		log.Fatal("Error loading env files")
	}
 
	fmt.Println("Port", port)
	
}