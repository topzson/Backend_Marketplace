package entity

import (
	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("mysql.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema

	database.AutoMigrate(&Role{}, &User{})

	db = database

	role1 := Role{
		Name: "Admin",
	}
	db.Model(&Role{}).Create(&role1)

	role2 := Role{
		Name: "Customer",
	}
	db.Model(&Role{}).Create(&role2)

}
