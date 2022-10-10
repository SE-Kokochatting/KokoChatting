package global

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

type Config struct{}

var globalConfig *Config
var once = new(sync.Once)

func init(){
	path,err := filepath.Abs("./config")
	if err != nil{
		panic(err)
	}
	viper.AddConfigPath(path)
	viper.SetConfigType("yaml")
	env := os.Getenv("ENVIRONMENT")
	switch env{
	case "production":
		viper.SetConfigName("prod")
	default:
		viper.SetConfigName("dev")
	}
	err = viper.ReadInConfig() //找到并读取配置文件
	if err != nil { // 捕获读取中遇到的error
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}

// GetGlobalConfig 单例模式获取全局配置
func GetGlobalConfig() *Config {
	once.Do(func(){
		globalConfig = new(Config)
	})
	return globalConfig
}

func (c *Config) GetConfigByName(configName string) interface{} {
	return viper.Get(configName)
}


func (c *Config) GetConfigByPath(path string) (string,error) {
	paths := strings.Split(path,".")
	v := c.GetConfigByName(paths[0])
	if str,ok := v.(string);ok{
		return str,nil
	}
	var i = 1
	for res,ok := v.(map[string]interface{});ok;i++{
		v = res[paths[i]]
		res,ok = v.(map[string]interface{})
	}
	switch v.(type){
	case int:
		return strconv.Itoa(v.(int)),nil
	case string:
		return v.(string),nil
	}
	return "",ConfigPathError
}
