package domain

import (
	"time"
)

type Contact struct {
	ID        string    `json:"id" gorm:"primaryKey;default:uuid_generate_v4()"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Subject   string    `json:"subject"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"` // ✅ Automatically managed
	UpdatedAt time.Time `json:"updated_at"` // ✅ Automatically managed
}
