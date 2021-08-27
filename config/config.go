package config

import (
	"log"

	"github.com/spf13/viper"
)

type App struct {
	ServerHost           string `mapstructure:"host"`
	ServerPort           string `mapstructure:"port"`
	KafkaBroker          string `mapstructure:"kafka_broker"`
	KafkaTopicDispatcher string `mapstructure:"kafka_topic_dispatcher"`
	RedisAdress          string `mapstructure:"redis_adress"`
	RedisPassword        string `mapstructure:"redis_password"`
	RedisDatabase        uint   `mapstructure:"redis_database"`
	GinMode              string `mapstructure:"gin_mode"`
	DatabaseUser         string `mapstructure:"db_user"`
	DatabasePass         string `mapstructure:"db_pass"`
	DatabaseSSL          string `mapstructure:"db_ssl"`
	DatabaseName         string `mapstructure:"db_name"`
	DatabaseHost         string `mapstructure:"db_host"`
	JwtSecretKey         string `mapstructure:"secret_key"`
}

var envConfig *viper.Viper
var conf *App

func Init() {

	envConfig = viper.New()
	envConfig.AddConfigPath(".")
	envConfig.AddConfigPath("../")
	envConfig.SetConfigType("env")
	envConfig.SetConfigName(`.env`)

	if err := envConfig.ReadInConfig(); err != nil {
		log.Fatalf("Error on reading the envConfig file: %v", err)
	}
	marshallErr := envConfig.Unmarshal(&conf)
	if marshallErr != nil {
		log.Fatalf("Error on unmarshalling the envConfig file: %v", marshallErr)
	}

}

func GetConfig() *App {
	return conf
}
