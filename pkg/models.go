package pkg

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	Id       uint64 `gorm:"primaryKey"`
	Username string
	Password string
	Email    string
	Phone    string
	Image    string
}

type Admin struct {
	gorm.Model
	Id     uint64 `gorm:"primaryKey"`
	User   User   `gorm:"foreignKey:UserId"`
	UserId uint64
}

type Task struct {
	gorm.Model
	Id          uint64 `gorm:"primaryKey"`
	UserId      uint64
	Name        string
	Description string
}

func init() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	DB = db

	if err != nil {
		panic("failed to connect database")
	}

	err = DB.AutoMigrate(&User{}, &Admin{}, &Task{})
	if err != nil {
		panic("Error autoMigrate: ")
	}
}
