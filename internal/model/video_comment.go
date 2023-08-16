package model

import (
    "time"
    "github.com/jinzhu/gorm"
	"errors"
)
type VideoComment struct {
    ID          uint32    `json:"id"`
    UserID      uint32    `json:"user_id"`
    VideoID     uint32    `json:"video_id"`
    Content     string    `json:"content"`
    CommentTime time.Time `json:"create_date"`
}

func (vc VideoComment) CommentAction(db *gorm.DB) (uint32, error) {
    result := db.Table("video_comments").Create(&vc)
    if result.Error != nil {
        return 0, result.Error
    }
    // 执行查询获取最后插入的自增主键值
    var comment VideoComment
	result = db.Table("video_comments").
		Where("user_id = ? AND video_id = ?", vc.UserID, vc.VideoID).
        Model(&vc).
		First(&comment)

    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
           return 0, result.Error
        }
        return 0, result.Error
    }
    
    return comment.ID, nil
}

func (vc VideoComment) CommentDelete(db *gorm.DB) error {
    result := db.Table("video_comments").Delete(&vc)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func (vc VideoComment) IsCommented(db *gorm.DB) (bool, error) {
    var count int64
    result := db.Table("video_comments").
        Model(&VideoComment{}).
        Where("user_id = ? AND id = ?", vc.UserID, vc.ID).
        Count(&count)
    if result.Error != nil {
        return false, result.Error
    }
    
    return count>0, nil
}

func (vc VideoComment) GetComment(db *gorm.DB) (*VideoComment, error) {
    var videoComment VideoComment

	result := db.Table("video_comments").
                Where("id = ?", vc.ID).First(&videoComment)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &videoComment, nil
}

func (vc VideoComment) CommentIdList(db *gorm.DB) ([]uint32, error) {
    var commentList []uint32
    result := db.Table("video_comments").
        Select("id").
        Where("video_id = ?", vc.VideoID).
        Pluck("video_id", &commentList)
    if result.RecordNotFound() {
    // 查询结果为空，可以根据需求进行相应的处理
        return nil, nil
    }
    if result.Error != nil {
        return nil, result.Error
    }
    
    return commentList, nil
}