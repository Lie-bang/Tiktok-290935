package Service

import (
	"context"
	"douyin/cmd/favorite/dal/db"
	"douyin/kitex_gen/douyinfavorite"
	"log"
)

type FavouriteCountUserService struct {
	ctx context.Context
}

func NewFavouriteCountUserService(ctx context.Context) *FavouriteCountUserService {
	return &FavouriteCountUserService{ctx: ctx}
}

func (f *FavouriteCountUserService) FavouriteCountUser(req *douyinfavorite.DouyinFavoriteCountUserRequest) (int64, error) {

	count, err := db.QueryFavoriteCountByUser(f.ctx, req)
	if err != nil {
		log.Print(err)
		return count, err
	}

	return count, err

}
