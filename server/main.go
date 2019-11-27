package main

import (
	"api-egressos/database"
	"api-egressos/handler"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

func main() {
	r := chi.NewRouter()

	database.CreateConnection()

	cors := cors.New(getCorsOptions())

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler)

	r.Get("/egressos", handler.GetProfiles)
	// r.Get("/", handler.GetPage)

	// abspath, _ := filepath.Abs("./pages/static")
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(abspath))))

	fmt.Println("Listening on port: 8000")

	http.ListenAndServe(":8000", r)
}

func getCorsOptions() cors.Options {
	return cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Äccept", "Authorization", "Content-Type", "X-XSRF-Token"},
		ExposedHeaders:   []string{"link"},
		AllowCredentials: true,
		MaxAge:           300,
	}
}
