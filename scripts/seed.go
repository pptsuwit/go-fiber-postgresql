package main

import (
	"fmt"
	"go-fiber-crud/app/config"
	"go-fiber-crud/app/model"
	"go-fiber-crud/scripts/mock"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	config.InitConfig()
	var err error
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Bangkok",
		viper.GetString("db.host"),
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.database"),
		viper.GetString("db.port"),
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}, &gorm.Config{
		// TranslateError: true,
		// Logger: new(model.SqlLogger),
		// DryRun: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	db.Migrator().DropTable(
		model.Customer{},
		model.User{},
		model.Asset{},
	)
	db.AutoMigrate(
		model.Customer{},
		model.User{},
		model.Asset{},
	)
	mock.SeedUser(db)
	mock.SeedCustomer(db)
	println("Seeded database")
}
