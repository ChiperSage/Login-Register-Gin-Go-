package models

import (
    "time"

    "github.com/jinzhu/gorm"
)

type User struct {
    UserID            uint      `gorm:"primary_key" json:"user_id"`
    Username          string    `gorm:"type:varchar(100);unique_index" json:"username"`
    Password          string    `json:"password"`
    LoginAttempts     int       `json:"login_attempts"`
    LastLoginAttempt  time.Time `json:"last_login_attempt"`
    RememberMeToken   string    `json:"remember_me_token"`
    CreatedAt         time.Time `json:"created_at"`
    UpdatedAt         time.Time `json:"updated_at"`
}

type Role struct {
    RoleID    uint      `gorm:"primary_key" json:"role_id"`
    RoleName  string    `gorm:"type:varchar(100);unique_index" json:"role_name"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

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
