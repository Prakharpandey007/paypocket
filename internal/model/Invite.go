package model
import (
	"github.com/google/uuid"
	"time"
)
type Invite struct{
    ID uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
   PasswordHash string `json:"-" gorm:"type:varchar(255);not null"` 
	Email       string `json:"email" gorm:"type:varchar(255);not null"`
	FirstName   string `json:"firstName" gorm:"type:varchar(255);not null"`
	LastName    string `json:"lastName" gorm:"type:varchar(255);not null"`
	PhoneNumber string `json:"phoneNumber" gorm:"type:varchar(255);not null"`
	InviteToken string `json:"inviteToken" gorm:"type:varchar(255);not null"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type AcceptInviteRequest struct {
	Password    string `json:"password" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	FirstName   string `json:"firstName" validate:"required"`
	LastName    string `json:"lastName" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	InviteToken string `json:"inviteToken" validate:"required"`
}
