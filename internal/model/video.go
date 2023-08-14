package model

import (
    "time"
    "github.com/jinzhu/gorm"
    "errors"
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

func (v Video) GetVideo(db *gorm.DB) (*Video, error) {
	var video Video

	// 使用 GORM 的 First 方法查询数据库中的记录
	result := db.Table("videos").Where("video_id = ?", v.VideoID).First(&video)

	// 检查查询时是否发生错误
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 如果没有找到匹配的记录，返回错误信息
			return nil, result.Error
		}
		// 如果发生其他错误，返回错误信息
		return nil, result.Error
	}

	// 如果找到了匹配的记录，返回查询到的用户和nil错误
	return &video, nil
}

func (v Video) GetVideoList(db *gorm.DB) ([]uint32, error) {
    var videoList []uint32
    result := db.Table("videos").
        Select("video_id").
        Where("user_id = ?", v.UserID).
        Pluck("video_id", &videoList)
    if result.RecordNotFound() {
    // 查询结果为空，可以根据需求进行相应的处理
        return nil, nil
    }
    if result.Error != nil {
        return nil, result.Error
    }
    
    return videoList, nil
}