package model

import (
    "time"
    "github.com/jinzhu/gorm"
)

type Video struct {
    VideoID uint32
    UserID uint32
    Title string
    FilePath string
    UploadTime time.Time
}

func (v Video) PublishAction(db *gorm.DB) error {
    if err := db.Table("videos").Create(&v).Error; err != nil {
		return err
	}
	return nil
}