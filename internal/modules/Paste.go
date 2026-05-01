package modules

import "time"

type Paste struct {
	ID       uint `gorm:"primaryKey"`
	Title    string
	Content  string
	IsTitled bool `gorm:"default:0"`

	Author string `gorm:"default:'Anonymous'"`

	TopicID uint
	Topic   Topic `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
