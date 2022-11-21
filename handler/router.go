package handler

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Router() {
	router := mux.NewRouter()
	router.HandleFunc("/tasktodo", GetTasks).Methods("GET")
	router.HandleFunc("/createtask", CreateTasks).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
