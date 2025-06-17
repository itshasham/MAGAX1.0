package models

import (
	"time"

	"github.com/google/uuid"
)

type Coupon struct {
	ID             uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Code           string     `gorm:"not null" json:"code"`
	IsActive       bool       `gorm:"default:true;not null" json:"is_active"`
	StartDate      *time.Time `json:"start_date,omitempty"`
	EndDate        *time.Time `json:"end_date,omitempty"`
	UsageLimit     *int       `json:"usage_limit,omitempty"`
	UsageCount     int        `gorm:"default:0" json:"usage_count"`
	MinOrderAmount *float64   `json:"min_order_amount,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}
