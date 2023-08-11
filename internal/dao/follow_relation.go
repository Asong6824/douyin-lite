package dao

import (
	"douyin/internal/model"
)

func (d *Dao) Follow(userid uint32, followerid uint32) error {
	fr := model.FollowRelation{
		UserID: userid,
    	FollowerID: followerid,
	}
	return fr.Follow(d.engine)
}

func (d *Dao) Unfollow(userid uint32, followerid uint32) error {
	fr := model.FollowRelation{
		UserID: userid,
    	FollowerID: followerid,
	}
	return fr.Unfollow(d.engine)
}

func (d *Dao) IsFollow(userid uint32, myuserid uint32) (bool, error) {
	fr := model.FollowRelation{
		UserID: userid,
    	FollowerID: myuserid,
	}
	return fr.IsFollow(d.engine)
}


func (d *Dao) FollowList(userid uint32) ([]uint32, error) {
	fr := model.FollowRelation{
    	FollowerID: userid,
	}
	return fr.FollowList(d.engine)
}