package routers

import (
	"context"
	"net/http"
	"github.com/go-chi/chi"
	"api/controllers"
)

func GetRouter() *chi.Mux {
	router := chi.NewRouter()
	
	router.Get("/", func(response http.ResponseWriter, request *http.Request) {
		response.Write([]byte("alive"))
	})
	
	router.Route("/books", func(router chi.Router) {
		router.Post("/", controllers.CreateBook)
		router.Get("/", controllers.RetrieveAllBooks)
		router.Route("/{id}", func(router chi.Router) {
			router.Use(GetIDContextFromURL)
			router.Get("/", controllers.RetrieveBookById)
			router.Put("/", controllers.UpdateBookById)
			router.Delete("/", controllers.DeleteBookById)
		})
	})

	return router
}

func GetIDContextFromURL(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		id := chi.URLParam(request, "id")
		idContext := context.WithValue(request.Context(), "id", id)
		nextHandler.ServeHTTP(response, request.WithContext(idContext))
	})
}