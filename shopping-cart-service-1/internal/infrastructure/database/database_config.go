package database

import (
	"log"

	"github.com/jinzhu/configor"
)

type DBConfig struct {
	Host     string `env:"DB_HOST" default:"localhost"`
	Port     string `env:"DB_PORT" default:"3306"`
	User     string `env:"DB_USER" default:"user123"`
	Password string `env:"DB_PASSWORD" default:"pass123"`
	Database string `env:"DB_DATABASE" default:"shopping_cart"`
	Type     string `env:"DB_TYPE" default:"mysql"`
}

func NewConfig() (conf DBConfig, err error) {
	log.Println("On loading database configuration")
	err = configor.New(&configor.Config{Environment: "development"}).Load(&conf)

	log.Println("configuration loaded: ")
	log.Println(conf)
	return
}
