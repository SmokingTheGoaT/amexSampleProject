package main

import (
	"amexSampleProject/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const defaultPort = "8080"

func main() {
	r := mux.NewRouter()
	handlers := handler.New(*r)
	handlers.CreateHandlers()

	log.Fatal(http.ListenAndServe(":"+defaultPort, r))
}
