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