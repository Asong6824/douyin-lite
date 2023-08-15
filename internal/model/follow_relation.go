package model

import (
	"github.com/jinzhu/gorm"
)

type FollowRelation struct {
    ID uint32
    UserID uint32
    FollowerID uint32
}

func (fr FollowRelation) Follow(db *gorm.DB) error {
    result := db.Table("follow_relations").Create(&fr)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func (fr FollowRelation) Unfollow(db *gorm.DB) error {
    result := db.Table("follow_relations").Delete(&fr)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func (fr FollowRelation) IsFollow(db *gorm.DB) (bool, error) {
    var count int64
    result := db.Table("follow_relations").
        Model(&FollowRelation{}).
        Where("follower_id = ? AND user_id = ?", fr.FollowerID, fr.UserID).
        Count(&count)
    if result.Error != nil {
        return false, result.Error
    }
    
    return count>0, nil
}

func (fr FollowRelation) FollowList(db *gorm.DB) ([]uint32, error) {
    var followList []uint32
    result := db.Table("follow_relations").
        Select("user_id").
        Where("follower_id = ?", fr.FollowerID).
        Pluck("user_id", &followList)
    if result.RecordNotFound() {
    // 查询结果为空，可以根据需求进行相应的处理
        return nil, nil
    }
    if result.Error != nil {
        return nil, result.Error
    }
    
    return followList, nil
}