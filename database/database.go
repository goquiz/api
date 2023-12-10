package database

import (
	"fmt"
	"github.com/goquiz/api/database/models"
	"github.com/goquiz/api/helpers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var Database *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", helpers.Env.Database.Username, helpers.Env.Database.Password, helpers.Env.Database.Host, helpers.Env.Database.Port, helpers.Env.Database.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to connect to the database: %v\n", err.Error()))
	}

	log.Println("successfully connected to the database!")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running migrations")
	err = db.AutoMigrate(
		&models.Answer{},
		&models.HostedQuiz{},
		&models.Question{},
		&models.Quiz{},
		&models.User{},
	)
	if err != nil {
		panic("Failed to migrate: " + err.Error())
	}

	Database = db
}
