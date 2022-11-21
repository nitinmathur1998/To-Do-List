package handler

import (
	"encoding/json"
	"example.com/mod/config"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

var Db *gorm.DB

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var listtodo []config.Todo
	Db.Find(&listtodo)
	err := json.NewEncoder(w).Encode(&listtodo)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var listtodo config.Todo
	err := json.NewDecoder(r.Body).Decode(&listtodo)
	if err != nil {
		log.Fatal(err)
	}

	createdTasks := Db.Create(&listtodo)
	err = createdTasks.Error
	if err != nil {
		err := json.NewEncoder(w).Encode(err) //better way to handle error
		if err != nil {
			return
		}
	} else {
		err := json.NewEncoder(w).Encode(&listtodo) //better way to handle error
		if err != nil {
			return
		}
	}
}
