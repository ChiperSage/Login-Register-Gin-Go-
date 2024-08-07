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
