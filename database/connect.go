package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/azizanhakim/items-api-fiber/config"
	"github.com/azizanhakim/items-api-fiber/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Declare the variable for the database
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Failed to parse port string")
	}

	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	fmt.Println(dsn)
	// Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	// Migrate the database
	err = DB.AutoMigrate(
		&model.Item{},
		&model.Sales{},
		&model.SalesItem{},
		&model.Type{},
		&model.User{},
	)

	if err != nil {
		fmt.Print(err)
		panic("Failed to auto migrate models")
	}

	fmt.Println("Database Migrated")

}
