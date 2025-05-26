package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/davidcl24/favourites_service/app/config"
	"github.com/davidcl24/favourites_service/app/handlers"
	"github.com/davidcl24/favourites_service/app/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var router *chi.Mux
var db *sql.DB

func init() {
	router = chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	dbConfig := *config.NewEnvDBConfig()

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password, dbConfig.Database)

	db, _ = sql.Open("postgres", connectionString)
	if err := db.Ping(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
}

func main() {
	router.Mount("/api/faavourites", favouritesRouters())
	http.ListenAndServe(":4000", router)
}

func favouritesRouters() chi.Router {
	router := chi.NewRouter()
	dbWrapper := models.NewDB(db)
	favouritesHandler := handlers.FavouriteHandler{DB: dbWrapper}

	router.Use(middleware.Logger)

	router.Get("/user/{user_id}", favouritesHandler.ListUserFavourites)
	router.Get("/{id}", favouritesHandler.GetFavourite)
	router.Post("/", favouritesHandler.CreateFavourite)
	router.Delete("/{id}", favouritesHandler.DeleteFavourite)
	router.Delete("/user/{user_id}", favouritesHandler.ClearUserFavourites)

	return router
}
