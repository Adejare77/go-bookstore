package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Adejare77/bookStore/internal/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	routes.BookStoreRoutes(router)
	fmt.Println("Startin Server on Port " + os.Getenv("PORT"))

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
