package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDb() {
	var err error
	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") + " password=" +
		os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT")
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database")
		os.Exit(1)
	}
}
