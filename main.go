package main

import (
	"example.com/mod/handler"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	dialect  = "postgres"
	host     = "localhost"
	dbPort   = "5432"
	user     = "nmathur"
	dbName   = "to_do"
	password = "abc123"
)

var db *gorm.DB
var err error

func main() {
	//database connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)

	//opening connection to database
	db, err = gorm.Open(dialect, dbURI)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to database")
	}

	//close connection when main function finishes
	//what is proper way to handle error here
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	//make migrations to the database if not been created
	db.AutoMigrate(&handler.Todo{})

	//API ROUTES
	router := mux.NewRouter()
	router.HandleFunc("/tasktodo", handler.GetTasks).Methods("GET")
	router.HandleFunc("/createtask", handler.CreateTasks).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))

}
