package helpers

import (
	"errors"
	"fmt"
	"hoxt/internal/modules"

	"gorm.io/gorm"
)

func CreatePasteIfTopicExists(db *gorm.DB, IdTopic uint, paste modules.Paste) (modules.Paste, error) {

	if err := db.Transaction(func(tx *gorm.DB) error {
		var topic modules.Topic
		// Check if topic exists
		if err := tx.Where("id = ?", IdTopic).First(&topic).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("topic with '%d' id does not exist", IdTopic)
			}
			return err
		}

		// Topic exists, associate it and create the post
		// paste.TopicID = topic.ID
		if err := tx.Create(&paste).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return modules.Paste{}, err
	}
	return paste, nil
}
