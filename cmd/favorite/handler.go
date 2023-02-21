package main

import (
	"context"
	Service "douyin/cmd/favorite/service"
	"douyin/kitex_gen/douyinfavorite"
	"fmt"
	"log"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, request *douyinfavorite.DouyinFavoriteActionRequest) (resp *douyinfavorite.DouyinFavoriteActionResponse, err error) {
	// TODO: Your code here...
	suc, err := Service.NewFavouriteActionService(ctx).FavouriteAction(request)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	resp = new(douyinfavorite.DouyinFavoriteActionResponse)
	resp.StatusCode = 0
	var StatusMsg string
	if suc {
		StatusMsg = "Favorite action success"
	} else {
		StatusMsg = "Something wrong happend when do Favorite action, please check the log."
	}

	resp.StatusMsg = &StatusMsg

	return resp, nil
}

// FavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, request *douyinfavorite.DouyinFavoriteListRequest) (resp *douyinfavorite.DouyinFavoriteListResponse, err error) {
	// TODO: Your code here...
	fmt.Println("get in handler FavoriteList")
	videoList, err := Service.NewFavouriteListService(ctx).FavouriteList(request)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	resp = new(douyinfavorite.DouyinFavoriteListResponse)
	resp.StatusCode = 0
	StatusMsg := "Get Favorite List Success"
	resp.StatusMsg = &StatusMsg
	resp.VideoList = videoList
	fmt.Println("get out handler FavoriteList")
	return resp, nil
}

// FavoriteJudge implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteJudge(ctx context.Context, request *douyinfavorite.DouyinFavoriteJudgeRequest) (resp *douyinfavorite.DouyinFavoriteJudgeResponse, err error) {
	// TODO: Your code here...
	res, err := Service.NewFavouriteJudgeService(ctx).FavouriteJudge(request)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	resp = new(douyinfavorite.DouyinFavoriteJudgeResponse)
	resp.StatusCode = 0
	resp.IfFav = res
	StatusMsg := "Favorite Judge Success"
	resp.StatusMsg = &StatusMsg

	return resp, err
}

// FavoriteCountUser implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteCountUser(ctx context.Context, request *douyinfavorite.DouyinFavoriteCountUserRequest) (resp *douyinfavorite.DouyinFavoriteCountUserResponse, err error) {
	// TODO: Your code here...
	fmt.Println("get in handler FavoriteCountUser")
	res, err := Service.NewFavouriteCountUserService(ctx).FavouriteCountUser(request)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	resp = new(douyinfavorite.DouyinFavoriteCountUserResponse)
	resp.FavoriteCount = res
	resp.StatusCode = 0
	StatusMsg := "Get Favorite Count By User Success"
	resp.StatusMsg = &StatusMsg
	fmt.Println("get out handler FavoriteCountUser")
	return resp, nil
}
