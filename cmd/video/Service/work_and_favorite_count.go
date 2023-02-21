package Service

import (
	"context"
	"douyin/cmd/video/dal/db"
	"douyin/kitex_gen/douyinvideo"
	"log"
)

type WorkAndFavoriteCountService struct {
	ctx context.Context
}

func NewWorkAndFavoriteCountService(ctx context.Context) *WorkAndFavoriteCountService {
	return &WorkAndFavoriteCountService{ctx: ctx}
}

func (f *WorkAndFavoriteCountService) WorkAndFavoriteCount(req *douyinvideo.Douyin_Work_And_Favorite_CountRequest) (int64, int64, error) {

	//err := db.UpdateCommentCount(f.ctx, req)

	//return nil
	workCount, favCount, err := db.QueryWorkAndFavoriteCount(f.ctx, req)
	if err != nil {
		log.Print(err)
		return 0, 0, err
	}
	return workCount, favCount, err
}
