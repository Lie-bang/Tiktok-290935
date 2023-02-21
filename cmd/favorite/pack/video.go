package pack

import (
	"douyin/kitex_gen/douyinfavorite"
	"douyin/kitex_gen/douyinvideo"
)

func DouyinVideoUserToDouyinFavoriteUser(user *douyinvideo.User) *douyinfavorite.User {
	return &douyinfavorite.User{
		Id:            (*user).Id,
		Name:          (*user).Name,
		FollowCount:   (*user).FollowerCount,
		FollowerCount: (*user).FollowerCount,
		IsFollow:      (*user).IsFollow,
	}
}

func DouyinVideoToDouyinFavoriteVideo(videos []*douyinvideo.Video) []*douyinfavorite.Video {
	var fVideo []*douyinfavorite.Video
	for _, v := range videos {

		fVideo = append(fVideo, &douyinfavorite.Video{
			Id:            (*v).Id,
			Author:        DouyinVideoUserToDouyinFavoriteUser(v.Author),
			PlayUrl:       (*v).PlayUrl,
			CoverUrl:      (*v).CoverUrl,
			FavoriteCount: (*v).FavoriteCount,
			CommentCount:  (*v).CommentCount,
			IsFavorite:    (*v).IsFavorite,
			Title:         (*v).Title,
		})
	}
	return fVideo
}
