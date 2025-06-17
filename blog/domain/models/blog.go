package models

import "time"

type Blog struct {
	ID              int       `gorm:"primaryKey;autoIncrement"`
	Title           string    `gorm:"type:text;not null"`
	Slug            string    `gorm:"type:text;not null;uniqueIndex"`
	Content         string    `gorm:"type:text;not null"`
	FeaturedImg     string    `gorm:"type:text"`
	Status          string    `gorm:"type:blog_status;not null"`   // ✅ now uses PostgreSQL enum
	Category        string    `gorm:"type:blog_category;not null"` // ✅
	Author          string    `gorm:"type:text;not null"`
	Views           int       `gorm:"default:0"`
	MetaTitle       string    `gorm:"type:text"`
	MetaDescription string    `gorm:"type:text"`
	MetaKeywords    string    `gorm:"type:text"`
	ReadTime        int       `gorm:"default:0"`
	SortOrder       int       `gorm:"default:0"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
}
