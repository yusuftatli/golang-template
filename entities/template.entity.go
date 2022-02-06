package entities

import "time"

type Template struct {
	ID        uint64    `gorm:"primaryKey"`
	Name      string    `gorm:"column:name"`
	Type      string    `gorm:"column:type"`
	Size      string    `gorm:"column:size"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
