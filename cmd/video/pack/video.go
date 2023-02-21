package pack

import (
	"context"
	"douyin/cmd/video/dal/db"
	"douyin/cmd/video/rpc"
	"douyin/kitex_gen/douyinfavorite"
	"douyin/kitex_gen/douyinvideo"
	"log"
)

func VideoDbToVideoService(dbList *[]db.Video_db, user []*douyinvideo.User, uid int64) []*douyinvideo.Video {

	var video []*douyinvideo.Video

	for k, v := range *dbList {

		if int64(v.UserId) == (*user[k]).Id {
			user[k].IsFollow = true
		}

		temp, err := rpc.FavoriteJudge(context.Background(), &douyinfavorite.DouyinFavoriteJudgeRequest{
			UserId:  uid,
			VideoId: int64(v.Id),
		})

		if err != nil {
			log.Print(err)
			return nil
		}

		var IsFav bool

		if temp.IfFav == 1 {
			IsFav = true
		} else {
			IsFav = false
		}

		t := douyinvideo.Video{
			Id:            int64(v.Id),
			Author:        user[k],
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: int64(v.FavouriteCount),
			CommentCount:  int64(v.CommentCount),
			IsFavorite:    IsFav,
			Title:         v.Title,
		}
		video = append(video, &t)
	}
	return video
}
