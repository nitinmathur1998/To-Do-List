package factory

import (
	"example.com/mod/config"
	"example.com/mod/handler"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var err error

func InitializeDb() {
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", config.Host, config.User, config.DbName, config.Password, config.DbPort)
	handler.Db, err = gorm.Open(config.Dialect, dbURI)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to database")
	}
	//make migrations to the database if not been created
	handler.Db.AutoMigrate(&config.Todo{})
}

func Close() {
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(handler.Db)
}
