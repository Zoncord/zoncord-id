package models

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MigrateModels(db *gorm.DB) {
	// user models
	db.AutoMigrate(&User{})
	// oauth2 models
	db.AutoMigrate(&Application{}, &Grant{}, &RefreshToken{})
}

func GetDSN() string {
	// get Data Source Name
	var dsn string
	dsn += "host=" + os.Getenv("SQL_HOST") + " "
	dsn += "user=" + os.Getenv("SQL_USER") + " "
	dsn += "password=" + os.Getenv("SQL_PASSWORD") + " "
	dsn += "dbname=" + os.Getenv("SQL_DATABASE") + " "
	dsn += "port=" + os.Getenv("SQL_PORT") + " "
	dsn += "sslmode=" + os.Getenv("SQL_SSL_MODE") + " "
	return dsn
}

func CreateAdmin(db *gorm.DB) {
	// Create admin user
	var admin User
	db.Where("email= ?", os.Getenv("ADMIN_EMAIL")).First(&admin)
	if admin.ID != 1 {
		admin = User{
			Email:    os.Getenv("ADMIN_EMAIL"),
			Password: os.Getenv("ADMIN_PASSWORD"),
		}
		db.Create(&admin)
		db.Save(&admin)
	}
	log.Printf("Admin created ID%d", admin.ID)
}

func CreateMasterApplication(db *gorm.DB) {
	// Create master application
	var masterApp Application
	db.Where("name = ?", "master").First(&masterApp)
	if masterApp.ID != 1 {
		masterApp = Application{
			Name:   "master",
			UserID: 1,
		}
		db.Create(&masterApp)
		db.Save(&masterApp)
	}
	log.Printf("Master application created ID%d", masterApp.ID)
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
	CreateAdmin(db)
	CreateMasterApplication(db)
	return db
}

func GetDB() *gorm.DB {
	db := InitDB()
	return db
}
