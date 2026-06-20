package main

import (
	"fmt"
	"os"
	"log"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main(){

	//loading all environmental variables
	godotenv.Load()

	//loading the dabase connection
	//DBURL=os.Getenv("DB_URL")

	fmt.Println("Hello, World!")

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT environment variable is not set") 
	}


	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	fmt.Println(port)
}