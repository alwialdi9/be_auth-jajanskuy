package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Username  string    `json:"username" gorm:"unique"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password" gorm:"not null"`
	FirstName string    `json:"first_name" gorm:"not null"`
	LastName  string    `json:"last_name" gorm:"not null"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt time.Time `json:"deleted_at,omitempty" gorm:"index"`
}

func (User) TableName() string {
	return "_users"
}

func CreateUser(db *gorm.DB, user *User) error {
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByID(db *gorm.DB, id int) (*User, error) {
	var user User
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByUsername(db *gorm.DB, username string) (*User, error) {
	var user User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(db *gorm.DB, email string) (*User, error) {
	var user User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
