package main

import (
	"api/controllers"
	_ "api/docs" // docs is generated by Swag CLI, you have to import it.
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.PathPrefix("/swagger").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
		httpSwagger.DeepLinking(false),
	)).Methods(http.MethodGet)

	apiRouter := router.PathPrefix("/api").Subrouter()

	apiRouter.HandleFunc("Cards/", controllers.GetCards).Methods(http.MethodGet)
	apiRouter.HandleFunc("Cards/{id}", controllers.GetCard).Methods(http.MethodGet)
	apiRouter.HandleFunc("Cards/", controllers.CreateCard).Methods(http.MethodPost)
	apiRouter.HandleFunc("Cards/{id}", controllers.UpdateCard).Methods(http.MethodPut)
	apiRouter.HandleFunc("Cards/{id}", controllers.DeleteCard).Methods(http.MethodDelete)

	err := http.ListenAndServe(":4000", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}