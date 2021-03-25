package main

import (
	"log"

	"./internal/test"
	// "github.com/jinzhu/gorm"
)

// var db *gorm.DB

func main() {
	log.Println("starting db...")
	test.GetDB()
	log.Println("wrapping up...")
}
