package util

import (
	"github.com/spf13/viper"
)
type Config struct {
    DBDriver      string `mapstructure:"DB_DRIVER"`
    DBSource      string `mapstructure:"DB_SOURCE"`
    HttpServerAddress string `mapstructure:"HTTP_SERVER_ADDR"`
}

func LoadConfig(path string) (config Config, err error) {
	//add file path
    viper.AddConfigPath(path)
    viper.SetConfigName("app")
    viper.SetConfigType("env")

	//read config file
	err=viper.ReadInConfig();
	if err!=nil{
		return
	}
	err=viper.Unmarshal(&config)
	return
}