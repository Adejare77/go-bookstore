package routes

import (
	"github.com/Adejare77/bookStore/internal/controllers"
	"github.com/gorilla/mux"
)

var BookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/books/{id}/", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/books", controllers.PostBook).Methods("POST")
	router.HandleFunc("/books/", controllers.PostBook).Methods("POST")
	router.HandleFunc("/books/{id}", controllers.UpdateBookById).Methods("PUT")
	router.HandleFunc("/books/{id}/", controllers.UpdateBookById).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.DeleteBookById).Methods("DELETE")
	router.HandleFunc("/books/{id}/", controllers.DeleteBookById).Methods("DELETE")
}
