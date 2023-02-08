package pack

import (
	"douyin/kitex_gen/douyinmessage"
	"douyin/kitex_gen/douyinrelation"
	"douyin/kitex_gen/douyinuser"
)

func UUserToRUser(usersMap map[int64]*douyinuser.User) []*douyinrelation.User {
	var rUsers []*douyinrelation.User
	for _, u := range usersMap {
		rUsers = append(rUsers, &douyinrelation.User{
			UserId:        u.UserId,
			Username:      u.Username,
			FollowCount:   u.FollowCount,
			FollowerCount: u.FollowerCount,
			IsFollow:      u.IsFollow,
			Avatar:        u.Avatar,
		})
	}
	return rUsers
}

func ToFirstMessage(messages []*douyinmessage.Message, userId int64) (map[int64]string, map[int64]int64) {
	contents := make(map[int64]string)
	msgTypes := make(map[int64]int64)
	for _, m := range messages {
		if m != nil {
			if userId == m.ToUserId {
				msgTypes[m.FromUserId] = 0
				contents[m.FromUserId] = m.Content
			} else {
				msgTypes[m.ToUserId] = 1
				contents[m.ToUserId] = m.Content
			}
		}
	}
	return contents, msgTypes
}

func ToFriendUser(usersMap map[int64]*douyinuser.User, contents map[int64]string, msgTypes map[int64]int64) []*douyinrelation.FriendUser {
	var rUsers []*douyinrelation.FriendUser
	for _, u := range usersMap {
		rUsers = append(rUsers, &douyinrelation.FriendUser{
			User: &douyinrelation.User{
				UserId:        u.UserId,
				Username:      u.Username,
				FollowCount:   u.FollowCount,
				FollowerCount: u.FollowerCount,
				IsFollow:      u.IsFollow,
				Avatar:        u.Avatar,
			},
			Message: contents[u.UserId],
			MsgType: msgTypes[u.UserId],
		})
	}
	return rUsers
}
