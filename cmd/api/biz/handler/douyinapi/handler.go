package douyinapi

import (
	"douyin/cmd/api/biz/model/douyinapi"
	"douyin/kitex_gen/douyinmessage"
	"douyin/kitex_gen/douyinrelation"
	"douyin/kitex_gen/douyinuser"
	"douyin/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type BaseResponse struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

type ApiUserResponse struct {
	StatusCode int64           `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string          `json:"status_msg"`  // 返回状态描述
	User       *douyinapi.User `json:"user"`        // 用户信息
}

type ApiUsersResponse struct {
	StatusCode int64             `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string            `json:"status_msg"`  // 返回状态描述
	UserList   []*douyinapi.User `json:"user_list"`   // 用户列表
}

type ApiFriendUserResponse struct {
	StatusCode int64                   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string                  `json:"status_msg"`  // 返回状态描述
	UserList   []*douyinapi.FriendUser `json:"user_list"`   // 用户列表
}

//type ApiMessageResponse struct {
//	StatusCode  int64                `json:"status_code"`  // 状态码，0-成功，其他值-失败
//	StatusMsg   string               `json:"status_msg"`   // 返回状态描述
//	MessageList []*douyinapi.Message `json:"message_list"` // 消息列表
//}

type ApiMessageResponse struct {
	StatusCode  int32       `json:"status_code"`  // 状态码，0-成功，其他值-失败
	StatusMsg   string      `json:"status_msg"`   // 返回状态描述
	MessageList interface{} `json:"message_list"` // 消息列表
	//MessageList interface{} `json:"message_list,omitempty"` // 消息列表
}

func UserUserToApiUser(user *douyinuser.User) *douyinapi.User {
	if user == nil {
		return nil
	}
	return &douyinapi.User{
		ID:              user.UserId,
		Name:            user.Username,
		FollowCount:     user.FollowCount,
		FollowerCount:   user.FollowerCount,
		IsFollow:        user.IsFollow,
		Avatar:          "http://192.168.0.109:8080/myavatar.jpg",
		BackgroundImage: "http://192.168.0.109:8080/wall.jpg",
		Signature:       "OH! MY GOD!",
		TotalFavorited:  1,
		WorkCount:       1,
		FavoriteCount:   1,
	}
}

func RelationUserToApiUser(user *douyinrelation.User) *douyinapi.User {
	if user == nil {
		return nil
	}
	return &douyinapi.User{
		ID:              user.UserId,
		Name:            user.Username,
		FollowCount:     user.FollowCount,
		FollowerCount:   user.FollowerCount,
		IsFollow:        user.IsFollow,
		Avatar:          "http://192.168.0.109:8080/myavatar.jpg",
		BackgroundImage: "http://192.168.0.109:8080/wall.jpg",
		Signature:       "OH! MY GOD!",
		TotalFavorited:  1,
		WorkCount:       1,
		FavoriteCount:   1,
	}
}

func RelationUsersToApiUsers(users []*douyinrelation.User) []*douyinapi.User {
	res := make([]*douyinapi.User, 0)
	for _, u := range users {
		if n := RelationUserToApiUser(u); n != nil {
			res = append(res, n)
		}
	}
	return res
}

func RelationFriendUserToApiFriendUser(user *douyinrelation.FriendUser) *douyinapi.FriendUser {
	if user == nil {
		return nil
	}
	return &douyinapi.FriendUser{
		ID:            user.User.UserId,
		Name:          user.User.Username,
		FollowCount:   user.User.FollowCount,
		FollowerCount: user.User.FollowerCount,
		IsFollow:      user.User.IsFollow,
		Message:       user.Message,
		MsgType:       user.MsgType,
		Avatar:        "http://192.168.0.109:8080/avatar.jpg",
	}
}

func RelationFriendUsersToApiFriendUsers(users []*douyinrelation.FriendUser) []*douyinapi.FriendUser {
	res := make([]*douyinapi.FriendUser, 0)
	for _, u := range users {
		if n := RelationFriendUserToApiFriendUser(u); n != nil {
			res = append(res, n)
		}
	}
	return res
}

func MsgMessageToApiMessage(msg *douyinmessage.Message) *douyinapi.Message {
	if msg == nil {
		return nil
	}
	return &douyinapi.Message{
		ID:         msg.MsgId,
		ToUserID:   msg.ToUserId,
		FromUserID: msg.FromUserId,
		Content:    msg.Content,
		CreateTime: msg.CreateTime,
	}
}

type TestMessage struct {
	ID         int64  `json:"id""`
	ToUserID   int64  `json:"to_user_id"`
	FromUserID int64  `json:"from_user_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

func MsgMessageToTestMessage(msg *douyinmessage.Message) *TestMessage {
	if msg == nil {
		return nil
	}
	return &TestMessage{
		ID:         msg.MsgId,
		ToUserID:   msg.ToUserId,
		FromUserID: msg.FromUserId,
		Content:    msg.Content,
		CreateTime: 1675843303941,
	}
}

func MsgMessagesToTestMessages(msgs []*douyinmessage.Message) []*TestMessage {
	res := make([]*TestMessage, 0)
	for _, m := range msgs {
		if n := MsgMessageToTestMessage(m); n != nil {
			res = append(res, n)
		}
	}
	return res
}

func MsgMessagesToApiMessages(msgs []*douyinmessage.Message) []*douyinapi.Message {
	res := make([]*douyinapi.Message, 0)
	for _, m := range msgs {
		if n := MsgMessageToApiMessage(m); n != nil {
			res = append(res, n)
		}
	}
	return res
}

func SendResponse(c *app.RequestContext, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, BaseResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
	})
}

func SendUserResponse(c *app.RequestContext, err error, user *douyinapi.User) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, ApiUserResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		User:       user,
	})
}

//type TestUser struct {
//	ID            int64  `form:"id" json:"id" query:"id"`
//	Name          string `form:"name" json:"name" query:"name"`
//	FollowCount   int64  `form:"follow_count" json:"follow_count" query:"follow_count"`
//	FollowerCount int64  `form:"follower_count" json:"follower_count" query:"follower_count"`
//	IsFollow      bool   `form:"is_follow" json:"is_follow" query:"is_follow"`
//	Avatar        string `form:"avatar" json:"avatar" query:"avatar"`
//	Work          int64  `json:"work_count" `
//	Like          int64  `json:"favorite_count" `
//}
//
//type ApiUserResponseTest struct {
//	StatusCode int64    `json:"status_code"` // 状态码，0-成功，其他值-失败
//	StatusMsg  string   `json:"status_msg"`  // 返回状态描述
//	User       TestUser `json:"user"`        // 用户信息
//}
//
//func SendUserResponseTest(c *app.RequestContext, err error, user *douyinapi.User) {
//	Err := errno.ConvertErr(err)
//	testUser := TestUser{
//		ID:            user.ID,
//		Name:          user.Name,
//		FollowCount:   user.FollowCount,
//		FollowerCount: user.FollowerCount,
//		IsFollow:      user.IsFollow,
//		Avatar:        user.Avatar,
//		Work:          12,
//		Like:          12,
//	}
//	c.JSON(consts.StatusOK, ApiUserResponseTest{
//		StatusCode: Err.ErrCode,
//		StatusMsg:  Err.ErrMsg,
//		User:       testUser,
//	})
//}

func SendUsersResponse(c *app.RequestContext, err error, users []*douyinapi.User) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, ApiUsersResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		UserList:   users,
	})
}

func SendFriendUsersResponse(c *app.RequestContext, err error, users []*douyinapi.FriendUser) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, ApiFriendUserResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		UserList:   users,
	})
}

func SendMessageTestResponse(c *app.RequestContext, err error, messages []*TestMessage) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, ApiMessageResponse{
		StatusCode:  int32(Err.ErrCode),
		StatusMsg:   Err.ErrMsg,
		MessageList: messages,
		//MessageList: messages,
	})
}

func SendMessageResponse(c *app.RequestContext, err error, messages []*douyinapi.Message) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, ApiMessageResponse{
		StatusCode:  int32(Err.ErrCode),
		StatusMsg:   Err.ErrMsg,
		MessageList: messages,
		//MessageList: messages,
	})
}
