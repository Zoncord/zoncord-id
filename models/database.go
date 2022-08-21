package models

import (
	"github.com/Zoncord/zoncord-id/models/oauth2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var dbase *gorm.DB

func MigrateModels(db *gorm.DB) {
	// user models
	db.AutoMigrate(&User{})
	// oauth2 models
	db.AutoMigrate(&oauth2.Application{}, &oauth2.Grant{}, &oauth2.RefreshToken{})
}

func GetDSN() string {
	// get Data Source Name
	var dsn string
	dsn += "host=" + os.Getenv("DB_HOST") + " "
	dsn += "user=" + os.Getenv("DB_USER") + " "
	dsn += "password=" + os.Getenv("DB_PASSWORD") + " "
	dsn += "dbname=" + os.Getenv("DB_NAME") + " "
	dsn += "port=" + os.Getenv("DB_PORT") + " "
	dsn += "sslmode=" + os.Getenv("DB_SSL_MODE") + " "
	return dsn
}

func InitDB() *gorm.DB {
	// Init database
	dsn := GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Миграция схем
	MigrateModels(db)
	return db
}

func GetDB() *gorm.DB {
	dbase = InitDB()
	return dbase
}
