package models

import "time"

type Board struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Uuid      string    `json:"uuid"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
