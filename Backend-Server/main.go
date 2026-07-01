package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
	"github.com/sanrose-bhets/Job-Grinder/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	//loading all environmental variables
	godotenv.Load()

	//port validation
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	//=========================================DATABASE CONNECTIONS================================================================

	//db connection validation
	DBURL := os.Getenv("DB_URL")
	if DBURL == "" {
		log.Fatal("DBURL environment variable is not set")
	}

	//connection to db
	conn, err := sql.Open("postgres", DBURL)
	if err != nil {
		log.Fatal("Can't connect to database:", err)
	}

	queries := database.New(conn)

	//pass this so that handlers have access to db
	apiCfg := apiConfig{
		DB: queries,
	}

	//==========================ROUTERS================================================================
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/health", readinessHandler)
	v1Router.Get("/err", errorHandler)

	v1Router.Post("/accounts", apiCfg.accountCreateHandler)
	v1Router.Post("/companies", apiCfg.companyCreateHandler)
	v1Router.Post("/interviews", apiCfg.interviewCreateHandler)
	v1Router.Post("/jobs", apiCfg.jobCreateHandler)
	v1Router.Post("/applications", apiCfg.applicationCreateHandler)

	//mounting the v1Router as to enable http request to server
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Printf("Server is Starting on Port: %v", port)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(port)
}
