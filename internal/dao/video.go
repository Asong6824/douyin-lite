package dao

import (
	"douyin/internal/model"
	"time"
)


func (d *Dao) PublishAction(userid uint32, title string, filepath string) error {
	video := model.Video{
		UserID: userid,
    	Title: title,
    	FilePath: filepath,
		UploadTime: time.Now(),
	}
	return video.PublishAction(d.engine)
}

//根据VideoId获取视频信息
func (d *Dao) GetVideo(videoid uint32) (*model.Video, error) {
	video := model.Video{
		VideoID: videoid,
	}
	return video.GetVideo(d.engine)
}

func (d *Dao) GetVideoList(userid uint32) ([]uint32, error) {
	video := model.Video{
		UserID: userid,
	}
	return video.GetVideoList(d.engine)
}

//ActionType 0 minus, 1 plus
func (d *Dao) ModifyVideoFavoriteCount(videoid uint32, ActionType int) error {
	video := model.Video{
		VideoID: videoid,
	}
	if ActionType == 1 {
		return video.PlusFavoriteCount(d.engine)
	} else {
		return video.MinusFavoriteCount(d.engine)
	}
}

func (d *Dao) ModifyCommentCount(videoid uint32, ActionType int) error {
	video := model.Video{
		VideoID: videoid,
	}
	if ActionType == 1 {
		return video.PlusCommentCount(d.engine)
	} else {
		return video.MinusCommentCount(d.engine)
	}
}