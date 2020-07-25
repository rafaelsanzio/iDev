package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"./services"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", services.GetInfoAPI).Methods("GET")
	router.HandleFunc("/server-info/{name}", services.GetServerInfo).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
