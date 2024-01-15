package config

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

}

func InitTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = ict
}
