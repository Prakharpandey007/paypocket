package model

import (
	"time"

	"github.com/google/uuid"
)

// Org is the DB model for organizations.
type Org struct {
	ID uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`

	Name         string `json:"name" gorm:"type:varchar(255);not null"`
	Industry     string `json:"industry" gorm:"type:varchar(255);not null"`
	Address      string `json:"address" gorm:"type:varchar(255);not null"`
	Country      string `json:"country" gorm:"type:varchar(100);not null"`
	City         string `json:"city" gorm:"type:varchar(100);not null"`
	Website      string `json:"website" gorm:"type:varchar(255)"`
	ContactEmail string `json:"contact_email" gorm:"type:varchar(255);not null"`
	ContactPhone string `json:"contact_phone" gorm:"type:varchar(50);not null"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateOrgRequest struct {
	Name         string `json:"name" validate:"required"`
	Industry     string `json:"industry" validate:"required"`
	Address      string `json:"address" validate:"required"`
	Country      string `json:"country" validate:"required"`
	City         string `json:"city" validate:"required"`
	Website      string `json:"website"`
	ContactEmail string `json:"contact_email" validate:"required,email"`
	ContactPhone string `json:"contact_phone" validate:"required,phone"`
}

type InviteEmployeeRequest struct {
	Email                   string    `json:"email" validate:"required,email"`
	Role                    string    `json:"role" validate:"required"`
	FirstName               string    `json:"first_name" validate:"required"`
	LastName                string    `json:"last_name" validate:"required"`
	Message                 string    `json:"message"`
	BaseSalary              float64   `json:"base_salary" validate:"required,numeric"`
	Department              string    `json:"department" validate:"required"`
	Bonus                   float64   `json:"bonus" validate:"numeric"`
	JoiningDate             time.Time `json:"joining_date" validate:"required,date"` // Format: "YYYY-MM-DD"
	EmploymentType          string    `json:"employment_type" validate:"required"`   // e.g., "Full-time", "Part-time", "Contract"
	OvertimeRate            float64   `json:"overtime_rate" validate:"numeric"`
	Allowances              float64   `json:"allowances" validate:"numeric"`
	HealthInsurance         float64   `json:"health_insurance" validate:"numeric"`
	RetirementBenefits      float64   `json:"retirement_benefits" validate:"numeric"`
	StockOptions            float64   `json:"stock_options" validate:"numeric"`
	StockOptionsVested      float64   `json:"stock_options_vested" validate:"numeric"`
	StockOptionsUnvested    float64   `json:"stock_options_unvested" validate:"numeric"`
	StockOptionsExpiration  time.Time `json:"stock_options_expiration"`
	StockOptionsGrantDate   time.Time `json:"stock_options_grant_date"`
	StockOptionsStrikePrice float64   `json:"stock_options_strike_price" validate:"numeric"`
	StockOptionsQuantity    int       `json:"stock_options_quantity" validate:"numeric"`
	StockOptionsType        string    `json:"stock_options_type"`                     // e.g., "ISO", "NSO"
	StockOptionsStatus      string    `json:"stock_options_status"`                   // e.g., "Active", "Vested", "Expired"
	Designation             string    `json:"designation" validate:"required"`        // e.g., "Software Engineer", "Product Manager"
	ReportingToID           *string   `json:"reporting_to_id"`                        // Optional, can be null if no direct manager
	PhoneNumber             string    `json:"phone_number" validate:"required,phone"` // Phone number of the employee
	
}
