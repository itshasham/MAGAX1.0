package models

import "time"

type FaqCategory string
type FaqStatus string

const (
	CategoryTerms   FaqCategory = "terms"
	CategoryPayment FaqCategory = "payment"
	CategoryContact FaqCategory = "contact"
	CategoryHelp    FaqCategory = "help"
	CategorySupport FaqCategory = "support"

	StatusEnabled  FaqStatus = "enabled"
	StatusDisabled FaqStatus = "disabled"
)

type Faq struct {
	ID        uint        `gorm:"primaryKey;autoIncrement"`
	Title     string      `gorm:"type:varchar(255)"`
	Content   string      `gorm:"type:text"`
	Category  FaqCategory `gorm:"type:faq_category"` // ← matches custom enum type
	status    FaqStatus   `gorm:"type:faq_status"`   // ← matches custom enum type
	SortOrder int
	CreatedAt time.Time
	UpdatedAt time.Time
}
