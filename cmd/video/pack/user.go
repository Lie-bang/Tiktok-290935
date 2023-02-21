package pack

import (
	"douyin/kitex_gen/douyinuser"
	"douyin/kitex_gen/douyinvideo"
	"douyin/pkg/consts"
)

func DyUserUserToDyVideoUser(user *douyinuser.User) *douyinvideo.User {
	ava := consts.AvatarAddr
	ba := consts.BackgroundImageAddr
	st := consts.Signature
	return &douyinvideo.User{
		Id:              user.UserId,
		Name:            user.Username,
		FollowCount:     &user.FollowCount,
		FollowerCount:   &user.FollowerCount,
		IsFollow:        user.IsFollow,
		Avatar:          &ava,
		TotalFavorited:  &user.TotalFavorited,
		WorkCount:       &user.WorkCount,
		FavoriteCount:   &user.FavoriteCount,
		BackgroundImage: &ba,
		Signature:       &st,
	}

}
