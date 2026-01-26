package model

import (
	"github.com/google/uuid"
	"time"
)
// User is the DB model that will be migrated into your local database. 
type User struct { 
	ID uuid.UUID `json:"id" gorm:"type:char(36);primaryKey" `
	Email string `json:"email" gorm:"type:varchar(255);uniqueIndex;not null" `
	PasswordHash string `json:"-" gorm:"type:varchar(255);not null"` 
	// RoleID uuid.UUID `json:"roleId" gorm:"type:char(36);not null" `
	FirstName string `json:"firstName" gorm:"type:varchar(100)" `
	LastName string `json:"lastName" gorm:"type:varchar(100)" `
	PhoneNumber string `json:"phoneNumber" gorm:"type:varchar(30)" `
	CreatedAt time.Time `json:"createdAt" `
	UpdatedAt time.Time `json:"updatedAt" `
}
type SignupRequest struct {
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	// RoleId      string `json:"roleId" binding:"required"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
}

type LoginRequest struct {
	Email string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}
