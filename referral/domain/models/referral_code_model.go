package models

import (
	"time"

	"github.com/google/uuid"
)

type ReferralCodeModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	Code      string    `gorm:"uniqueIndex;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ReferralCode struct {
	ID     uuid.UUID
	UserID uuid.UUID
	Code   string
}

func (ReferralCodeModel) TableName() string {
	return "referral_codes"
}
