package pack

import (
	"douyin/cmd/user/dal/db"
	"douyin/kitex_gen/douyinrelation"
	"douyin/kitex_gen/douyinuser"
)

func UserG(u *db.User, relationInfo *douyinrelation.User, favoriteCount, workCount, totalFavorite int64) *douyinuser.User {
	if u == nil {
		return nil
	}

	return &douyinuser.User{
		UserId:         int64(u.ID),
		Username:       u.Username,
		FollowCount:    relationInfo.FollowCount,
		FollowerCount:  relationInfo.FollowerCount,
		IsFollow:       relationInfo.IsFollow,
		Avatar:         "http://192.168.0.109:8080/avatar.jpg",
		TotalFavorited: totalFavorite,
		WorkCount:      workCount,
		FavoriteCount:  favoriteCount,
	}
}

func User(u *db.User, relationInfo *douyinrelation.User) *douyinuser.User {
	if u == nil {
		return nil
	}

	return &douyinuser.User{
		UserId:         int64(u.ID),
		Username:       u.Username,
		FollowCount:    relationInfo.FollowCount,
		FollowerCount:  relationInfo.FollowerCount,
		IsFollow:       relationInfo.IsFollow,
		Avatar:         "http://192.168.0.109:8080/avatar.jpg",
		TotalFavorited: 0,
		WorkCount:      0,
		FavoriteCount:  0,
	}
}

// Users pack list of user info
func Users(us []*db.User, relationInfos []*douyinrelation.User) []*douyinuser.User {
	users := make([]*douyinuser.User, 0)
	for idx, u := range us {
		if temp := User(u, relationInfos[idx]); temp != nil {
			users = append(users, temp)
		}
	}
	return users
}
