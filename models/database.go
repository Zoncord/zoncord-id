package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbase *gorm.DB

func InitDB() *gorm.DB {
	// Init database
	dsn := "host=localhost user=ZoncordAdmin password=ASF0o7z5 dbname=ZoncordID port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Миграция схем
	db.AutoMigrate(&User{})
	return db
}

func GetDB() *gorm.DB {
	dbase = InitDB()
	return dbase
}
