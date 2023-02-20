package db

import (
	"context"
	"douyin/cmd/video/dal/db"
	"douyin/kitex_gen/douyinfavorite"
	"douyin/pkg/consts"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

type Favorite_db struct {
	UserId       int64
	VideoId      int64
	FavoriteTime int64 // 在展示点赞列表时，倒序展示
}

func (f *Favorite_db) TableName() string {
	return consts.FavoriteTableName
}

func UpdateFavourite(ctx context.Context, request *douyinfavorite.DouyinFavoriteActionRequest, userid int64) error {
	var fav Favorite_db
	var video db.Video_db

	//需要更新两个表，更新videob表中视频点赞的数量，g更新favourite报表中user_id以及video_id的关系
	if request.ActionType == 1 { //点赞
		fav.UserId = userid
		fav.VideoId = request.VideoId
		res := DB.WithContext(ctx).Where("user_id = ? AND video_id = ?", fav.UserId, fav.VideoId).Find(&fav)
		if res.RowsAffected > 0 {
			log.Print("already liked!!!!")
			return nil
		}
		fav.UserId = userid
		fav.VideoId = request.VideoId
		fav.FavoriteTime = time.Now().Unix()
		resultFavorite := DB.WithContext(ctx).Create(&fav)

		if resultFavorite.Error != nil {
			log.Print(resultFavorite.Error)
			return resultFavorite.Error
		}

		fav.UserId = userid
		fav.VideoId = request.VideoId
		result := DB.WithContext(ctx).Model(&video).Where("id = ?", request.VideoId).
			Update("favourite_count", gorm.Expr("favourite_count + ?", 1))
		if result.Error != nil {
			log.Print(result.Error)
			return result.Error
		}

	} else if request.ActionType == 2 { //取消点赞
		fav.UserId = userid
		fav.VideoId = request.VideoId
		res := DB.WithContext(ctx).Where("user_id = ? ", fav.UserId).Find(&fav)
		if res.RowsAffected == 0 {
			log.Print("can't cancel like!!!!")
			return nil
		}
		fav.UserId = userid
		fav.VideoId = request.VideoId
		fmt.Println("cancel liked user_id: ", fav.UserId)
		fmt.Println("cancel liked video_id: ", fav.VideoId)
		resultFavorite := DB.WithContext(ctx).Where("user_id = ? AND video_id = ?", fav.UserId, fav.VideoId).Delete(&fav)

		if resultFavorite.Error != nil {
			log.Print(resultFavorite.Error)
			return resultFavorite.Error
		}
		fav.UserId = userid
		fav.VideoId = request.VideoId
		result := DB.WithContext(ctx).Model(&video).Where("id = ?", request.VideoId).
			Update("favourite_count", gorm.Expr("favourite_count - ?", 1))
		if result.Error != nil {
			log.Print(result.Error)
			return result.Error
		}
		//fav.UserId = userid
		//fav.VideoId = request.VideoId

	} else {
		log.Print("wrong action type")
		return nil
	}
	return nil
}

func QueryFavoriteVideoByUserId(ctx context.Context, request *douyinfavorite.DouyinFavoriteListRequest) (*[]int64, error) {
	var fav *[]Favorite_db
	fmt.Println("in QueryFavoriteVideoByUserId user_id :", request.UserId)
	result := DB.WithContext(ctx).Where("user_id = ?", request.UserId).Order("favorite_time desc").Find(&fav)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Println("fav now:", *fav)
	fmt.Println(len(*fav))
	var videoIds []int64
	for k, v := range *fav {
		fmt.Println("key now:", k)
		fmt.Println("videoID now:", v.VideoId)
		videoIds = append(videoIds, v.VideoId)
	}
	fmt.Println("get out for range")
	return &videoIds, nil
}

func QueryFavoriteJudge(ctx context.Context, request *douyinfavorite.DouyinFavoriteJudgeRequest) (int32, error) {
	var fav Favorite_db

	result := DB.WithContext(ctx).Where("user_id = ? AND video_id = ?", request.UserId, request.VideoId).Find(&fav)

	if result.RowsAffected == 0 {
		return 0, result.Error
	} else {
		return 1, result.Error
	}

}

func QueryFavoriteCountByUser(ctx context.Context, request *douyinfavorite.DouyinFavoriteCountUserRequest) (int64, error) {
	var fav []Favorite_db
	result := DB.WithContext(ctx).Where("user_id = ?", request.UserId).Find(&fav)
	if result.Error != nil {
		log.Print(result.Error)
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
