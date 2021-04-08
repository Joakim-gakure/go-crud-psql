package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitializeDB() *gorm.DB {
	db, err := gorm.Open("postgres", "user=manosriram password=password dbname=manosriram sslmode=disable")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("DB connected!")
	}
	return db
}
