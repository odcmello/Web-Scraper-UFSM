package main

import (
	"api-egressos/database"
	"api-egressos/handler"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	database.CreateConnection()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/egressos", handler.GetProfiles)

	fmt.Println("Listening on port: 8000")

	http.ListenAndServe(":8000", r)
}
