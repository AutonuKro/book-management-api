package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AutonuKro/go-book-management-api/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/", router)
	fmt.Println("server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
