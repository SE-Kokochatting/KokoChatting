package provider

import (
	"KokoChatting/global"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
)

func MysqlInit() (*gorm.DB, error){
	// use viper to read configuration
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	path = path + "/../config"
	config := viper.New()
	config.AddConfigPath(path)
	config.SetConfigName("dev")
	config.SetConfigType("yaml")
	if err := config.ReadInConfig(); err != nil {
		global.Logger.Error("read config error", zap.Error(err))
		// global.Logger.Panic("read config error", err)
		panic(err)
	}

	host := config.Get("mysql.host").(string)
	username := config.Get("mysql.username").(string)
	password := config.Get("mysql.password").(string)
	database := config.Get("mysql.database").(string)

	dns := username + ":" + password + "@(" + host + ")/" + database + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dns)
	return db, err
}