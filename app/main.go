package main

import (
	"fmt"
	"go-fiber-crud/app/config"
	"go-fiber-crud/app/config/database"
	"go-fiber-crud/app/config/logs"
	"go-fiber-crud/app/router"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	config.InitConfig()
	config.InitTimeZone()
	logs.InitLogs()
	db = database.New(db)

	// gin.SetMode(viper.GetString("app.mode"))
	router := router.InitRouter(db)
	router.Listen(fmt.Sprintf(":%v", viper.GetString("app.port")))
}
