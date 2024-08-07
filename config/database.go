package config

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "gin-user-management/models"
)

var DB *gorm.DB

func ConnectDatabase() {
    var err error
    DB, err = gorm.Open("mysql", "yourusername:yourpassword@tcp(localhost:3306)/yourdbname?charset=utf8mb4&parseTime=True&loc=Local")
    if err != nil {
        panic("Failed to connect to database!")
    }

    DB.AutoMigrate(&models.User{}, &models.Role{}, &models.Group{})
}
