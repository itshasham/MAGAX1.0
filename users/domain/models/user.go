package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID
	Wallet         string `gorm:"unique"`
	ReferralBy     *string
	Name           *string
	Email          *string
	PhoneNumber    *string
	SocialX        *string
	SocialTelegram *string
	Bonus          float64
	Amount         float64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func UUIDFromString(id string) uuid.UUID {
	uid, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil
	}
	return uid
}
