package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model        // Ini akan otomatis menambahkan Id, CreatedAt, UpdatedAt, dan DeletedAt
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now().UTC()
	return
}

// TableName overrides
func (User) TableName() string {
	return "users"
}
