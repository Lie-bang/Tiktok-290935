package pack

import (
	"context"
	"douyin/cmd/video/dal/db"
	"douyin/cmd/video/rpc"
	"douyin/kitex_gen/douyinfavorite"
	"douyin/kitex_gen/douyinvideo"
	"fmt"
	"log"
)

/*
进行db类型->kitex类型的转换
*/

/*
type Video(douyinvideo) struct {
	Id            int64  `thrift:"id,1,required" frugal:"1,required,i64" json:"id"`
	Author        *User  `thrift:"author,2,required" frugal:"2,required,User" json:"author"`
	PlayUrl       string `thrift:"play_url,3,required" frugal:"3,required,string" json:"play_url"`
	CoverUrl      string `thrift:"cover_url,4,required" frugal:"4,required,string" json:"cover_url"`
	FavoriteCount int64  `thrift:"favorite_count,5,required" frugal:"5,required,i64" json:"favorite_count"`
	CommentCount  int64  `thrift:"comment_count,6,required" frugal:"6,required,i64" json:"comment_count"`
	IsFavorite    bool   `thrift:"is_favorite,7,required" frugal:"7,required,bool" json:"is_favorite"`
	Title         string `thrift:"title,8,required" frugal:"8,required,string" json:"title"`
}
*/

/*
type Video_db struct {
	Id             int
	UserId         int
	PlayUrl        string
	CoverUrl       string
	FavouriteCount int
	CommentCount   int
	IsFavourite    int
	Title          string
	PubTime        int64
}
*/

/*
struct User {
  1: required i64 id  // 用户id
  2: required string name  // 用户名称
  3: optional i64 follow_count // 关注总数
  4: optional i64 follower_count  // 粉丝总数
  5: required bool is_follow  // true-已关注，false-未关注
}
*/

func VideoDbToVideoService(dbList *[]db.Video_db, user []*douyinvideo.User, uid int64) []*douyinvideo.Video {

	var video []*douyinvideo.Video

	fmt.Println("VideoDbToVideoService user:", user)

	fmt.Println("get in VideoDbToVideoService")
	for k, v := range *dbList {

		if int64(v.UserId) == (*user[k]).Id {
			user[k].IsFollow = true
		}

		fmt.Println("user now:", *(user[k]))

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
		//fmt.Println("uid now: ", uid)
		//fmt.Println("v.Id now: ", int64(v.Id))
		////针对自己刷到自己视频的情况做处理
		//if uid == int64(v.Id) {
		//	IsFav = true
		//}

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
