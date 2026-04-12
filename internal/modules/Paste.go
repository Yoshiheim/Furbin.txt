package modules

import "time"

type Paste struct {
	ID       uint
	Title    string `gorm:"default:'[NULL]'"`
	Content  string `gorm:"default:'[NULL]'"`
	IsTitled bool   `gorm:"default:0"`

	Author string `gorm:"default:'Anonymous'"`

	TopicID uint
	Topic   Topic `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
