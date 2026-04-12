package modules

import "time"

type Topic struct {
	ID          uint
	Name        string `gorm:"unique"`
	Description string
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
