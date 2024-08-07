package models

import (
    "time"

    "github.com/jinzhu/gorm"
)

type Group struct {
    ID        uint      `gorm:"primary_key" json:"id"`
    RoleID    uint      `json:"role_id"`
    UserID    uint      `json:"user_id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    Role      Role      `gorm:"foreignkey:RoleID"`
    User      User      `gorm:"foreignkey:UserID"`
}

func Migrate(db *gorm.DB) {
    db.AutoMigrate(&User{})
    db.AutoMigrate(&Role{})
    db.AutoMigrate(&Group{})
}
