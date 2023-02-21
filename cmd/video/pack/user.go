package pack

import (
	"douyin/kitex_gen/douyinuser"
	"douyin/kitex_gen/douyinvideo"
)

func DyUserUserToDyVideoUser(user *douyinuser.User) *douyinvideo.User {
	return &douyinvideo.User{
		Id:             user.UserId,
		Name:           user.Username,
		FollowCount:    &user.FollowCount,
		FollowerCount:  &user.FollowerCount,
		IsFollow:       user.IsFollow,
		Avatar:         &user.Avatar,
		TotalFavorited: &user.TotalFavorited,
		WorkCount:      &user.WorkCount,
		FavoriteCount:  &user.FavoriteCount,
	}

}
