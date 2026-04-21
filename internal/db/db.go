package db

import (
	"hoxt/data"
	"hoxt/internal/modules"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDataBase() {
	if data.Configs.DBFilename == "" {
		data.Configs.DBFilename = "data.db"
	}
	var err error
	DB, err = gorm.Open(sqlite.Open(data.Configs.DBFilename), &gorm.Config{
		TranslateError: true, // Позволяет GORM возвращать понятные ошибки
		Logger:         logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database")
	}

	DB.Exec("PRAGMA foreign_keys = ON")
	err = DB.AutoMigrate(&modules.Paste{}, &modules.Topic{})
	if err != nil {
		panic(err.Error())
	}

	var count int64
	DB.Model(&modules.Topic{}).Count(&count)

	if count >= 0 {

		cfg, err := data.LoadConfig("./data/config.json")
		if err != nil {
			log.Fatalln(err)
		}

		if len(cfg.Topics) > 0 {
			for _, t := range cfg.Topics {
				if t.Name == "" || t.Description == "" {
					log.Fatalf("We have some null data, bro: T:%s D:%s", t.Name, t.Description)
				}

				var topic modules.Topic
				result := DB.Where("name = ?", t.Name).First(&topic)

				if result.Error != nil {
					// не найден — создаём
					DB.Create(&modules.Topic{
						Name:        t.Name,
						Description: t.Description,
					})
				} else {
					// найден — обновляем
					topic.Description = t.Description
					DB.Save(&topic)
				}
			}
		}
		if len(cfg.Pastes) > 0 {
			for _, p := range cfg.Pastes {
				var paste modules.Paste

				result := DB.Where("title = ?", p.Title).First(&paste)

				if result.Error != nil {
					DB.Create(&modules.Paste{
						Title:    p.Title,
						Content:  p.Content,
						Author:   "WEBSITE SYSTEM",
						IsTitled: p.IsTitled,
						TopicID:  p.TopicIndex,
					})
				} else {
					paste.Content = p.Content
					paste.TopicID = p.TopicIndex
					DB.Save(&paste)
				}
			}
		}
	}
}
