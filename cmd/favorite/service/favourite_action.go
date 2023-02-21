package Service

import (
	"context"
	"douyin/cmd/favorite/dal/db"
	"douyin/kitex_gen/douyinfavorite"
	"log"
	"strconv"
)

type FavouriteActionService struct {
	ctx context.Context
}

func NewFavouriteActionService(ctx context.Context) *FavouriteActionService {
	return &FavouriteActionService{ctx: ctx}
}

func (f *FavouriteActionService) FavouriteAction(req *douyinfavorite.DouyinFavoriteActionRequest) (bool, error) {
	//操作数据库
	userid, err := strconv.ParseInt(req.Token, 10, 64)
	if err != nil {
		log.Print(err)
		return false, err
	}
	success, err := db.UpdateFavourite(f.ctx, req, userid)
	if err != nil {
		log.Print(err)
		return false, err
	}
	return success, nil
}
