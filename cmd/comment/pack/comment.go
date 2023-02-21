package pack

import (
	"context"
	"douyin/cmd/comment/dal/db"
	"douyin/cmd/comment/rpc"
	"douyin/kitex_gen/douyincomment"
	"douyin/kitex_gen/douyinuser"
)

func CommentDbToDouyinComment(cmt *db.CommentDb) (*douyincomment.Comment, error) {

	uUser, err := rpc.GerUser(context.Background(), &douyinuser.GetUserRequest{
		UserId:   cmt.UserId,
		ToUserId: cmt.UserId,
	})
	if err != nil {
		return nil, err
	}

	return &douyincomment.Comment{
		Id:         cmt.CommentId,
		User:       DouyinUserUserToDouyinCommentUser(uUser),
		Content:    cmt.Content,
		CreateDate: cmt.CreateDate,
	}, nil
}

func DouyinUserUserToDouyinCommentUser(user *douyinuser.User) *douyincomment.User {
	return &douyincomment.User{
		Id:            user.UserId,
		Name:          user.Username,
		FollowCount:   &user.FollowerCount,
		FollowerCount: &user.FollowerCount,
		IsFollow:      user.IsFollow,
	}
}
