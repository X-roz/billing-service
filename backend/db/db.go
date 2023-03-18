package db

import (
	"billing-service/internals/model"
	"flag"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func ConnectToDb() {
	// var (
	// 	DB_PATH string = os.Getenv("DB_PATH")
	// )

	Db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal()
	}

	migrateflag := flag.Bool("migrate", true, "Select the migrateDb is need to happen or not")
	flag.Parse()
	if *migrateflag {
		Db.AutoMigrate(&model.ItemDetails{})
	}
}

func DbManager() *gorm.DB {
	return Db
}