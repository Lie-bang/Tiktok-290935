package pack

import (
	"douyin/cmd/relation/dal/rdb"
	"douyin/kitex_gen/douyinmessage"
	"douyin/kitex_gen/douyinrelation"
)

func ToRUser(userRModels []*rdb.User, usersMap map[int64]string) []*douyinrelation.User {
	var rUsers []*douyinrelation.User
	for _, u := range userRModels {
		rUsers = append(rUsers, &douyinrelation.User{
			UserId:        u.ID,
			Username:      usersMap[u.ID],
			FollowCount:   u.FollowCount,
			FollowerCount: u.FollowerCount,
			IsFollow:      u.IsFollow,
			Avatar:        "empty",
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

func ToFriendUser(userRModels []*rdb.User, usersMap map[int64]string, contents map[int64]string, msgTypes map[int64]int64) []*douyinrelation.FriendUser {
	var rUsers []*douyinrelation.FriendUser
	for _, u := range userRModels {
		rUsers = append(rUsers, &douyinrelation.FriendUser{
			User: &douyinrelation.User{
				UserId:        u.ID,
				Username:      usersMap[u.ID],
				FollowCount:   u.FollowCount,
				FollowerCount: u.FollowerCount,
				IsFollow:      u.IsFollow,
				Avatar:        "test",
			},
			Message: contents[u.ID],
			MsgType: msgTypes[u.ID],
		})
	}
	return rUsers
}
