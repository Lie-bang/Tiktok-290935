package Service

import (
	"context"
	"douyin/cmd/favorite/dal/db"
	"douyin/cmd/favorite/pack"
	"douyin/cmd/favorite/rpc"
	"douyin/kitex_gen/douyinfavorite"
	"douyin/kitex_gen/douyinvideo"
)

type FavouriteListService struct {
	ctx context.Context
}

func NewFavouriteListService(ctx context.Context) *FavouriteListService {
	return &FavouriteListService{ctx: ctx}
}

func (f *FavouriteListService) FavouriteList(req *douyinfavorite.DouyinFavoriteListRequest) ([]*douyinfavorite.Video, error) {

	videoIds, err := db.QueryFavoriteVideoByUserId(f.ctx, req)
	if err != nil {
		return nil, err
	}
	videoList, err := rpc.GetFavoriteList(context.Background(), &douyinvideo.DouyinFavoriteListRequest{
		VideoId: *videoIds,
		UserId:  req.UserId,
	})
	if err != nil {
		return nil, err
	}
	fVideo := pack.DouyinVideoToDouyinFavoriteVideo(videoList)
	//给回最终的结果

	return fVideo, nil
}
