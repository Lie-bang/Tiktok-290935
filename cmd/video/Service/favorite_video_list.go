package Service

import (
	"context"
	"douyin/cmd/video/dal/db"
	"douyin/cmd/video/pack"
	"douyin/kitex_gen/douyinvideo"
	"log"
)

type FavoriteVideoListService struct {
	ctx context.Context
}

func NewFavoriteVideoListService(ctx context.Context) *FavoriteVideoListService {
	return &FavoriteVideoListService{ctx: ctx}
}

func (f *FavoriteVideoListService) FavoriteVideoList(req *douyinvideo.DouyinFavoriteListRequest) ([]*douyinvideo.Video, error) {
	res, err := db.QueryFavoriteVideoList(f.ctx, req.VideoId)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	userList := GetUserListFromVideoDb(res)
	douyinvideoUserList, err := GetUserListForVideo(userList, req.UserId)

	newVideo := pack.VideoDbToVideoService(res, douyinvideoUserList, req.UserId)
	return newVideo, nil

}
