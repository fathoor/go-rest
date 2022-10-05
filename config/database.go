package config

import (
	"go-rest/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	DSN := "host=" + Config.DB_HOST + " user=" + Config.DB_USER + " password=" + Config.DB_PASS +
		" dbname=" + Config.DB_NAME + " port=" + Config.DB_PORT + " sslmode=" + Config.DB_SSL_MODE

	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&models.Order{}, &models.Item{})
}

func FetchDB() *gorm.DB {
	return DB
}
