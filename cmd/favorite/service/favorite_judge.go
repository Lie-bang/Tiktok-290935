package Service

import (
	"context"
	"douyin/cmd/favorite/dal/db"
	"douyin/kitex_gen/douyinfavorite"
	"log"
)

type FavouriteJudgeService struct {
	ctx context.Context
}

func NewFavouriteJudgeService(ctx context.Context) *FavouriteJudgeService {
	return &FavouriteJudgeService{ctx: ctx}
}

func (f *FavouriteJudgeService) FavouriteJudge(req *douyinfavorite.DouyinFavoriteJudgeRequest) (int32, error) {
	res, err := db.QueryFavoriteJudge(f.ctx, req)

	if err != nil {
		log.Print(err)
		return res, err
	}

	return res, err

}
