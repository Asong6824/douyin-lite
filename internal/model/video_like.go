package model

import (
	"github.com/jinzhu/gorm"
)

type VideoLike struct {
    ID uint32
    UserID uint32
    VideoID uint32
}

func (vl VideoLike) Favorite(db *gorm.DB) error {
    result := db.Table("video_likes").Create(&vl)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func (vl VideoLike) Unfavorite(db *gorm.DB) error {
    result := db.Table("video_likes").Delete(&vl)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func (vl VideoLike) IsFavorite(db *gorm.DB) (bool, error) {
    var count int64
    result := db.Table("video_likes").
        Model(&VideoLike{}).
        Where("user_id = ? AND video_id = ?", vl.UserID, vl.VideoID).
        Count(&count)
    if result.Error != nil {
        return false, result.Error
    }
    
    return count>0, nil
}

func (vl VideoLike) FavoriteList(db *gorm.DB) ([]uint32, error) {
    var favoriteList []uint32
    result := db.Table("video_likes").
        Select("video_id").
        Where("user_id = ?", vl.UserID).
        Pluck("video_id", &favoriteList)
    if result.RecordNotFound() {
    // 查询结果为空，可以根据需求进行相应的处理
        return nil, nil
    }
    if result.Error != nil {
        return nil, result.Error
    }
    
    return favoriteList, nil
}