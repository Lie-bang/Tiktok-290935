package pack

import (
	"douyin/cmd/user/dal/db"
	"douyin/kitex_gen/douyinrelation"
	"douyin/kitex_gen/douyinuser"
)

func User(u *db.User, relationInfo *douyinrelation.User) *douyinuser.User {
	if u == nil {
		return nil
	}

	return &douyinuser.User{
		Id:            int64(u.ID),
		Name:          u.Username,
		FollowCount:   relationInfo.FollowCount,
		FollowerCount: relationInfo.FollowerCount,
		IsFollow:      relationInfo.IsFollow,
		Avatar:        nil,
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
