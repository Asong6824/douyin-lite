package dao

import (
	"douyin/internal/model"
	"time"
)

func (d *Dao) CommentAction(userId uint32, videoId uint32, content string, commentTime time.Time) (uint32, error) {
	vc := model.VideoComment{
		VideoID: videoId,
		UserID: userId,
		Content: content,
		CommentTime: commentTime,
	}
	return vc.CommentAction(d.engine)
}

func (d *Dao) CommentDelete(commentId uint32) error {
	vc := model.VideoComment{
		ID: commentId,
	}
	return vc.CommentDelete(d.engine)
}

func (d *Dao) IsCommented(commentId uint32, userId uint32) (bool, error) {
	vc := model.VideoComment{
		ID: commentId,
		UserID: userId,
	}
	return vc.IsCommented(d.engine)
}

func (d *Dao) GetComment(commentId uint32) (*model.VideoComment, error) {
	vc := model.VideoComment{
		ID: commentId,
	}
	return vc.GetComment(d.engine)
}

func (d *Dao) CommentIdList(videoId uint32) ([]uint32, error) {
	vc := model.VideoComment{
		VideoID: videoId,
	}
	return vc.CommentIdList(d.engine)
}