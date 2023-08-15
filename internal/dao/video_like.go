package dao

import (
	"douyin/internal/model"
)

func (d *Dao) Favorite(userid uint32, videoid uint32) error {
	vl := model.VideoLike{
		UserID: userid,
		VideoID: videoid,
	}
	return vl.Favorite(d.engine)
}

func (d *Dao) Unfavorite(userid uint32, videoid uint32) error {
	vl := model.VideoLike{
		UserID: userid,
		VideoID: videoid,
	}
	return vl.Unfavorite(d.engine)
}

func (d *Dao) IsFavorite(userid uint32, videoid uint32) (bool, error) {
	vl := model.VideoLike{
		UserID: userid,
		VideoID: videoid,
	}
	return vl.IsFavorite(d.engine)
}

func (d *Dao) FavoriteList(userid uint32) ([]uint32, error) {
	vl := model.VideoLike{
    	UserID: userid,
	}
	return vl.FavoriteList(d.engine)
}