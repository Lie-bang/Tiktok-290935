package Service

import (
	"context"
	"douyin/cmd/favorite/dal/db"
	"douyin/cmd/favorite/pack"
	"douyin/cmd/favorite/rpc"
	"douyin/kitex_gen/douyinfavorite"
	"douyin/kitex_gen/douyinvideo"
	"fmt"
)

type FavouriteListService struct {
	ctx context.Context
}

func NewFavouriteListService(ctx context.Context) *FavouriteListService {
	return &FavouriteListService{ctx: ctx}
}

func (f *FavouriteListService) FavouriteList(req *douyinfavorite.DouyinFavoriteListRequest) ([]*douyinfavorite.Video, error) {
	//操作数据库
	//res, err := db.QueryFavouriteList(f.ctx, req)
	//
	//if err != nil {
	//	log.Print(err)
	//	return nil, err
	//}
	////需要将(user_id, video_id) -> (video)
	//newRes, err := db.QueryVideoByFavourite(f.ctx, res)
	//fmt.Println("finish QueryVideoByFavourite")
	//if err != nil {
	//	log.Print(err)
	//	return nil, err
	//}
	//fmt.Println("get in Video_dbToVideo_Service")
	//newVideo := pack.Video_dbToVideo_Service(newRes)
	//fmt.Println("get out Video_dbToVideo_Service")
	//return newVideo, nil
	//先查数据库，给到全部的video_id
	fmt.Println("get in db.QueryFavoriteVideoByUserId")
	videoIds, err := db.QueryFavoriteVideoByUserId(f.ctx, req)
	fmt.Println("videoIds:", videoIds)
	if err != nil {
		return nil, err
	}
	fmt.Println("get out db.QueryFavoriteVideoByUserId")
	fmt.Println("get in rpc.GetFavoriteList")
	videoList, err := rpc.GetFavoriteList(context.Background(), &douyinvideo.DouyinFavoriteListRequest{
		VideoId: *videoIds,
		UserId:  req.UserId,
	})
	fmt.Println("get out rpc.GetFavoriteList")
	if err != nil {
		return nil, err
	}
	fmt.Println("packing")
	fVideo := pack.DouyinVideoToDouyinFavoriteVideo(videoList)
	fmt.Println("packing over")
	//给回最终的结果

	return fVideo, nil
}
