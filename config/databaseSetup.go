package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       *string `gorm:"not null;uniqueIndex:book_unique;type:varchar(60)"`
	Author      *string `gorm:"not null;uniqueIndex:book_unique;type:varchar(60)"`
	Publication *string `gorm:"not null;uniqueIndex:book_unique;type:varchar(60)"`
}

var db *gorm.DB

func Connect() {
	// load from .env if it exists, else use system env
	_ = godotenv.Overload()

	user := os.Getenv("USER")
	pwd := os.Getenv("MYSQL_PASSWORD")
	database := os.Getenv("DATABASE")

	// Data Source Name
	dsn := user + ":" + pwd + "@/" + database + "?charset=utf8&parseTime=True&loc=Local"

	d, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		log.Fatalf("Could not connect to %s Database. Ensure %s Exists", database, database)
	}

	db = d
}

func BookTable() {
	book := &Book{}
	db.AutoMigrate(&book)
}

func GetDB() *gorm.DB {
	return db
}
