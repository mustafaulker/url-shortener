package models

import "time"

type ShortURL struct {
	Code      string `gorm:"primaryKey;size:10"`
	FullURL   string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Expiry    time.Duration
	Clicks    uint
}
