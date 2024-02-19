package database

import (
	"fmt"
	"go-fiber-crud/app/model"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *gorm.DB {
	var err error
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Bangkok",
		viper.GetString("db.host"),
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.database"),
		viper.GetString("db.port"),
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}, &gorm.Config{
		TranslateError: true,
		// Logger: new(model.SqlLogger),
		// DryRun: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(
		model.Customer{},
		model.User{},
		model.Asset{},
	)
	return db
}
