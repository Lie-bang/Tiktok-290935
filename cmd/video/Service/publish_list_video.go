package Service

import (
	"context"
	"douyin/cmd/video/dal/db"
	"douyin/cmd/video/pack"
	"douyin/kitex_gen/douyinvideo"
	"log"
)

type PublishListVideoService struct {
	ctx context.Context
}

func NewPublishListVideoService(ctx context.Context) *PublishListVideoService {
	return &PublishListVideoService{ctx: ctx}
}

func (p *PublishListVideoService) PublishListVideo(req *douyinvideo.DouyinPublishListRequest) ([]*douyinvideo.Video, error) {

	res, err := db.QueryListVideo(p.ctx, int(req.UserId))
	if err != nil {
		log.Print(err)
		return nil, err
	}
	userList := GetUserListFromVideoDb(res)
	douyinvideoUserList, err := GetUserListForVideo(userList, req.UserId)

	newVideo := pack.VideoDbToVideoService(res, douyinvideoUserList, req.UserId)
	return newVideo, nil

}
