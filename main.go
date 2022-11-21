package main

import (
	"example.com/mod/factory"
	"example.com/mod/handler"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	factory.InitializeDb()
	handler.Router()
	factory.Close()
}
