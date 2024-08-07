package models

import (
    "time"

    "github.com/jinzhu/gorm"
)

type Role struct {
    RoleID    uint      `gorm:"primary_key" json:"role_id"`
    RoleName  string    `gorm:"type:varchar(100);unique_index" json:"role_name"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
