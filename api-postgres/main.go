package main

import (
	"api-postgres/configs"
	"api-postgres/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}
	r := chi.NewRouter()
	r.Post("/todos/", handlers.Create)
	r.Put("/todos/", handlers.Update)
	r.Delete("/todos/{id}", handlers.Delete)
	r.Get("/todos/", handlers.List)
	r.Get("/todos/{id}", handlers.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
