// Code generated by hertz generator.

package douyinapi

import (
	"context"
	"douyin/cmd/api/biz/model/douyinapi"
	"douyin/cmd/api/biz/mw"
	"douyin/cmd/api/biz/rpc"
	"douyin/kitex_gen/douyinmessage"
	"douyin/kitex_gen/douyinrelation"
	"douyin/kitex_gen/douyinuser"
	"douyin/pkg/consts"
	"douyin/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
)

// CreateUser .
// @router /douyin/user/register/ [POST]
func CreateUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyinapi.CreateUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}

	err = rpc.CreateUser(context.Background(), &douyinuser.CreateUserRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}
	mw.JwtMiddleware.LoginHandler(ctx, c)
}

// CheckUser .
// @router /douyin/user/login/ [POST]
func CheckUser(ctx context.Context, c *app.RequestContext) {
	mw.JwtMiddleware.LoginHandler(ctx, c)
}

// GetUser .
// @router /douyin/publish/action/ [POST]
func GetUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyinapi.GetUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendUserResponse(c, errno.ConvertErr(err), nil)
		return
	}

	//var userId int64
	//if req.Token != "" {
	//	JWTToken, err := mw.JwtMiddleware.ParseTokenString(req.Token)
	//	if err != nil {
	//		SendUserResponse(c, errno.ConvertErr(err), nil)
	//	}
	//	claims := jwt.ExtractClaimsFromToken(JWTToken)
	//	userId=int64(claims[consts.IdentityKey].(float64))
	//	println(userId)
	//}

	v, _ := c.Get(consts.IdentityKey)
	uUser, err := rpc.GerUser(context.Background(), &douyinuser.GetUserRequest{
		UserId:   v.(*douyinapi.User).ID,
		ToUserId: req.UserID,
	})
	if err != nil {
		SendUserResponse(c, errno.ConvertErr(err), nil)
		return
	}

	user := UserUserToApiUser(uUser)
	SendUserResponse(c, errno.Success, user)
}

// Action .
// @router /douyin/relation/action/ [POST]
func Action(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyinapi.ActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}

	v, _ := c.Get(consts.IdentityKey)
	err = rpc.Action(context.Background(), &douyinrelation.ActionRequest{
		UserId:     v.(*douyinapi.User).ID,
		ToUserId:   req.ToUserID,
		ActionType: req.ActionType,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}

	SendResponse(c, errno.Success)
}

// FollowList .
// @router /douyin/relation/follow/list/ [GET]
func FollowList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyinapi.FollowListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendUsersResponse(c, errno.ConvertErr(err), nil)
		return
	}

	v, _ := c.Get(consts.IdentityKey)
	rUsers, err := rpc.FollowList(context.Background(), &douyinrelation.FollowListRequest{
		UserId:   v.(*douyinapi.User).ID,
		ToUserId: req.UserID,
	})
	if err != nil {
		SendUsersResponse(c, errno.ConvertErr(err), nil)
		return
	}

	users := RelationUsersToApiUsers(rUsers)
	SendUsersResponse(c, errno.Success, users)
}

// FollowerList .
// @router /douyin/relation/follower/list/ [GET]
func FollowerList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyinapi.FollowerListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendUsersResponse(c, errno.ConvertErr(err), nil)
		return
	}

	v, _ := c.Get(consts.IdentityKey)
	rUsers, err := rpc.FollowerList(context.Background(), &douyinrelation.FollowerListRequest{
		UserId:   v.(*douyinapi.User).ID,
		ToUserId: req.UserID,
	})
	if err != nil {
		SendUsersResponse(c, errno.ConvertErr(err), nil)
		return
	}

	users := RelationUsersToApiUsers(rUsers)
	SendUsersResponse(c, errno.Success, users)
}

// FriendList .
// @router /douyin/relation/friend/list/ [GET]
func FriendList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyinapi.FriendListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendFriendUsersResponse(c, errno.ConvertErr(err), nil)
		return
	}

	v, _ := c.Get(consts.IdentityKey)
	rFriendUsers, err := rpc.FriendList(context.Background(), &douyinrelation.FriendListRequest{
		UserId:   v.(*douyinapi.User).ID,
		ToUserId: req.UserID,
	})
	if err != nil {
		SendFriendUsersResponse(c, errno.ConvertErr(err), nil)
		return
	}

	friendUsers := RelationFriendUsersToApiFriendUsers(rFriendUsers)
	SendFriendUsersResponse(c, errno.Success, friendUsers)
}

// ChatRecord .
// @router /douyin/message/action/ [POST]
func ChatRecord(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyinapi.ChatRecordRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendMessageTestResponse(c, errno.ConvertErr(err), nil)
		return
	}
	v, _ := c.Get(consts.IdentityKey)
	mChatRecord, err := rpc.ChatRecord(context.Background(), &douyinmessage.ChatRecordRequest{
		UserId:   v.(*douyinapi.User).ID,
		ToUserId: req.ToUserID,
	})
	if err != nil {
		SendMessageTestResponse(c, errno.ConvertErr(err), nil)
		return
	}

	chatRecord := MsgMessagesToTestMessages(mChatRecord)
	SendMessageTestResponse(c, errno.Success, chatRecord)
}

// SendMessage .
// @router /douyin/message/chat/ [GET]
func SendMessage(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyinapi.SendMessageRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}

	v, _ := c.Get(consts.IdentityKey)
	err = rpc.SendMessage(context.Background(), &douyinmessage.SendMessageRequest{
		UserId:     v.(*douyinapi.User).ID,
		ToUserId:   req.ToUserID,
		Content:    req.Content,
		ActionType: req.ActionType,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}

	SendResponse(c, errno.Success)
}
