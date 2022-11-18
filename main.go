package main

import (
	"encoding/json"
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

type Todo struct {
	Task    string `json:"task"`
	Checked bool   `json:"checked"`
}

var (
	toDoList = &Todo{Task: "Learn Go", Checked: true}
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
	defer db.Close()

	//make migrations to the database if not been created
	db.AutoMigrate(&Todo{})

	//API ROUTES
	router := mux.NewRouter()
	router.HandleFunc("/tasktodo", getTasks).Methods("GET")
	router.HandleFunc("/createtask", createTasks).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))

}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var listtodo []Todo
	db.Find(&listtodo)
	json.NewEncoder(w).Encode(&listtodo)
}

func createTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var listtodo Todo
	err := json.NewDecoder(r.Body).Decode(&listtodo)
	if err != nil {
		return
	}

	createdTasks := db.Create(&listtodo)
	err = createdTasks.Error
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(&listtodo)
	}
}
