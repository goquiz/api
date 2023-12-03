package database

import (
	"fmt"
	"github.com/bndrmrtn/goquiz_api/database/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var Database *gorm.DB

func Connect() {
	dsn := "root:@tcp(127.0.0.1:3306)/go_quiz?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to connect to the database: %v\n", err.Error()))
	}

	log.Println("successfully connected to the database!")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running migrations")
	err = db.AutoMigrate(
		&models.Question{},
		&models.Quiz{},
		&models.User{},
	)
	if err != nil {
		panic("Failed to migrate: " + err.Error())
	}

	Database = db
}
