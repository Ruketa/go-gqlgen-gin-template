package db

import (
	model "github.com/go-gqlgen-gin-template/graph/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Host     string
	Database string
	User     string
	Port     int
	Password string
}

var DB *gorm.DB

func LoadConfig() *ConfigList {
	config, err := ini.Load("./db/dbconfig.ini")
	if err != nil {
		panic(err)
	}
	return &ConfigList{
		Host:     config.Section("db").Key("host").String(),
		Database: config.Section("db").Key("database").String(),
		User:     config.Section("db").Key("user").String(),
		Port:     config.Section("db").Key("port").MustInt(),
		Password: config.Section("db").Key("password").String(),
	}
}

func InitDB() {

	config := LoadConfig()

	dns := "user=" + config.User + " password=" + config.Password + " host=" + config.Host + " dbname=" + config.Database + " sslmode=disable"
	gormDB, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = gormDB

	DB.AutoMigrate(&model.User{}, &model.Todo{})
}

func GetDB() *gorm.DB {
	return DB
}

func CloseDB() {
	database, err := DB.DB()
	if err != nil {
		panic(err)
	}
	database.Close()
}
