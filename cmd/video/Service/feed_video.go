package Service

import (
	"context"
	"douyin/cmd/video/dal/db"
	"douyin/cmd/video/pack"
	"douyin/cmd/video/rpc"
	"douyin/kitex_gen/douyinuser"
	"douyin/kitex_gen/douyinvideo"
	"fmt"
	"strconv"
	"time"
)

type FeedVideoService struct {
	ctx context.Context
}

func NewFeedVideoService(ctx context.Context) *FeedVideoService {
	return &FeedVideoService{ctx: ctx}
}

func (f *FeedVideoService) FeedVideo(req *douyinvideo.DouyinFeedRequest) ([]*douyinvideo.Video, int64, error) {

	res, err := db.QueryVideo(f.ctx, *req.LastestTime)
	if err != nil {
		fmt.Println("db.QueryVideo something wrong")
		return nil, 0, err
	}

	if len(*res) == 0 {
		return nil, time.Now().Unix(), nil
	}

	respLatesttime := (*res)[len(*res)-1].PubTime

	/*
		otel-collector-config.yaml.从video数据库表中进行video相关记录的查询
		2.根据返回的video_db类型数据调用其他rpc服务，得到想要的数据。
		3.将相关数据进行组装。
	*/

	userList := GetUserListFromVideoDb(res)

	uid, err := strconv.ParseInt(*req.Token, 10, 64)

	douyinvideoUserList, err := GetUserListForVideo(userList, uid)

	newVideo := pack.VideoDbToVideoService(res, douyinvideoUserList, uid)

	return newVideo, respLatesttime, nil
}

// GetUserListFromVideoDb 整合数据库中所包含的user_id
func GetUserListFromVideoDb(dbList *[]db.Video_db) *[]int64 {
	var UserIdList []int64
	for _, v := range *dbList {
		UserIdList = append(UserIdList, int64(v.UserId))

	}
	return &UserIdList
}

// GetUserListForVideo 调用rpc服务，返回user_id对应的Author的结构体信息
func GetUserListForVideo(userIdList *[]int64, userid int64) ([]*douyinvideo.User, error) {

	var UserModelList []*douyinvideo.User

	if userid != -1 {
		for _, v := range *userIdList {
			uUser, err := rpc.GerUser(context.Background(), &douyinuser.GetUserRequest{
				UserId:   userid,
				ToUserId: v,
			})

			if err != nil {
				return nil, err
			}
			user := pack.DyUserUserToDyVideoUser(uUser)
			UserModelList = append(UserModelList, user)
		}
		return UserModelList, nil
	} else {
		for _, v := range *userIdList {
			uUser, err := rpc.GerUser(context.Background(), &douyinuser.GetUserRequest{
				UserId:   v,
				ToUserId: v,
			})

			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			user := pack.DyUserUserToDyVideoUser(uUser)
			UserModelList = append(UserModelList, user)
		}
		return UserModelList, nil
	}

}
