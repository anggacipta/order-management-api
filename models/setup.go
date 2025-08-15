package models

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("order.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke database!", err)
	}
	// Migrasi model
	database.AutoMigrate(&User{}, &Product{}, &Order{}, &OrderItem{})
	DB = database
}

func SetupTestDB() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect test database")
	}
	db.AutoMigrate(&User{}, &Product{}, &Order{}, &OrderItem{})
	DB = db
}
