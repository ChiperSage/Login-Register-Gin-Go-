package config

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "gin-user-management/models"
)

var DB *gorm.DB

func ConnectDatabase() {
    var err error
    DB, err = gorm.Open("postgres", "host=localhost user=yourusername dbname=yourdbname sslmode=disable password=yourpassword")
    if err != nil {
        panic("Failed to connect to database!")
    }

    models.Migrate(DB)
}
