package config

import (
	"log"
	"github.com/spf13/viper"
)

// global var to access from any package
var LocalConfig *Config 

type Config struct{
	DBUser string `mapstructure:"DBUSER"`
	DBPass string `mapstructure:"DBPASS"`
	DBIp string `mapstructure:"DBIP"`
	DBName string `mapstructure:"DBNAME"`
	Port string `mapstructure:"PORT"`
}

// initializing config vars 
func InitConfig() *Config{
	viper.AddConfigPath(("."))
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	// automatically reads the config vars
	if err := viper.ReadInConfig(); err != nil{
		log.Fatal("Error reading env file", err)
	}
	var config *Config
	// converts the read config vars into mapped struct type
	if err := viper.Unmarshal(&config); err != nil{
		log.Fatal("Error reading env file", err)
	}
	return config
}

// calling InitConfig() to initialize the config vars
func SetConfig(){
	LocalConfig = InitConfig()
}