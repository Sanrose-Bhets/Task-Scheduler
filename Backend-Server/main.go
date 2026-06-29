package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
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


	v1Router:=chi.NewRouter()

	v1Router.HandleFunc("/health",readinessHandler)
	v1Router.HandleFunc("/err",errorHandler)


	//mounting the v1Router as to enable http request to server
	router.Mount("/v1", v1Router)



	srv:= &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}


	log.Printf("Server is Starting on Port: %v", port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(port)
}