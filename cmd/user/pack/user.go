package pack

import (
	"douyin/cmd/user/dal/db"
	"douyin/kitex_gen/douyinuser"
)

func User(u *db.User, followCount, followerCount int64, isFollow bool) *douyinuser.User {
	if u == nil {
		return nil
	}

	return &douyinuser.User{
		UserId:        int64(u.ID),
		Username:      u.Username,
		FollowCount:   followCount,
		FollowerCount: followerCount,
		IsFollow:      isFollow,
		Avatar:        "test",
	}
}

// Users pack list of user info
func Users(us []*db.User, fs, fers []int64, isfs []bool) []*douyinuser.User {
	users := make([]*douyinuser.User, 0)
	for idx, u := range us {
		if temp := User(u, fs[idx], fers[idx], isfs[idx]); temp != nil {
			users = append(users, temp)
		}
	}
	return users
}
