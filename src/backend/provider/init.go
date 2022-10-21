package provider

import (
	"KokoChatting/global"
	"github.com/jinzhu/gorm"
)

func MysqlInit() (*gorm.DB, error){
	// use viper to read configuration
	config := global.GetGlobalConfig()

	host := config.GetConfigByName("mysql.host").(string)
	username := config.GetConfigByName("mysql.username").(string)
	password := config.GetConfigByName("mysql.password").(string)
	database := config.GetConfigByName("mysql.database").(string)

	dns := username + ":" + password + "@(" + host + ")/" + database + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dns)
	return db, err
}