package main

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi"
)

func main(){

	//loading all environmental variables
	godotenv.Load()

	//loading the dabase connection
	DBURL=os.Getenv("DB_URL")

	fmt.Println("Hello, World!")

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT environment variable is not set") 
	}


	router := chi.NewRouter()

	fmt.Println(port)
}