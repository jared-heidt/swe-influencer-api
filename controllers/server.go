package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var router *mux.Router

func initHandlers() {
	router.HandleFunc("/api/creators", GetAllCreators).Methods("GET")
	router.HandleFunc("/api/creators/{id}", GetCreator).Methods("GET")
	router.HandleFunc("/api/creators", CreateCreator).Methods("POST")
	router.HandleFunc("/api/creators", UpdateCreator).Methods("PUT")
	router.HandleFunc("/api/creators/{id}", DeleteCreator).Methods("DELETE")
}

func Start() {
	router = mux.NewRouter()

	initHandlers()

	err := godotenv.Load("config/.env")

	if err != nil {
		fmt.Println("Error loading .env file: ", err.Error())
		return
	}
	port := os.Getenv("API_PORT")
	fmt.Println("Router initialized and listening on ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
