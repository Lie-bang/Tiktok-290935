package db

import (
	"context"
	"douyin/cmd/video/dal/db"
	"douyin/kitex_gen/douyinfavorite"
	"douyin/pkg/consts"
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

func UpdateFavourite(ctx context.Context, request *douyinfavorite.DouyinFavoriteActionRequest, userid int64) (bool, error) {
	var fav Favorite_db
	var video db.Video_db

	videoExist := DB.WithContext(ctx).Where("id = ?", request.VideoId).Find(&video)
	if videoExist.Error != nil {
		log.Print(videoExist.Error)
		return false, videoExist.Error
	}

	if videoExist.RowsAffected == 0 {
		log.Print("video_id doesn't exist!")
		return false, nil
	}

	//需要更新两个表，更新videob表中视频点赞的数量，更新favourite报表中user_id以及video_id的关系
	if request.ActionType == 1 { //点赞
		fav.UserId = userid
		fav.VideoId = request.VideoId
		res := DB.WithContext(ctx).Where("user_id = ? AND video_id = ?", fav.UserId, fav.VideoId).Find(&fav)
		if res.RowsAffected > 0 {
			log.Print("already liked!!!!")
			return false, nil
		}
		fav.UserId = userid
		fav.VideoId = request.VideoId
		fav.FavoriteTime = time.Now().Unix()
		resultFavorite := DB.WithContext(ctx).Create(&fav)

		if resultFavorite.Error != nil {
			log.Print(resultFavorite.Error)
			return false, resultFavorite.Error
		}

		fav.UserId = userid
		fav.VideoId = request.VideoId
		result := DB.WithContext(ctx).Model(&video).Where("id = ?", request.VideoId).
			Update("favourite_count", gorm.Expr("favourite_count + ?", 1))
		if result.Error != nil {
			log.Print(result.Error)
			return false, result.Error
		}

	} else if request.ActionType == 2 { //取消点赞
		fav.UserId = userid
		fav.VideoId = request.VideoId
		res := DB.WithContext(ctx).Where("user_id = ? ", fav.UserId).Find(&fav)
		if res.RowsAffected == 0 {
			log.Print("can't cancel like!!!!")
			return false, nil
		}
		fav.UserId = userid
		fav.VideoId = request.VideoId
		//fmt.Println("cancel liked user_id: ", fav.UserId)
		//fmt.Println("cancel liked video_id: ", fav.VideoId)
		resultFavorite := DB.WithContext(ctx).Where("user_id = ? AND video_id = ?", fav.UserId, fav.VideoId).Delete(&fav)

		if resultFavorite.Error != nil {
			log.Print(resultFavorite.Error)
			return false, resultFavorite.Error
		}
		fav.UserId = userid
		fav.VideoId = request.VideoId
		result := DB.WithContext(ctx).Model(&video).Where("id = ?", request.VideoId).
			Update("favourite_count", gorm.Expr("favourite_count - ?", 1))
		if result.Error != nil {
			log.Print(result.Error)
			return false, result.Error
		}

	} else {
		log.Print("wrong action type")
		return false, nil
	}
	return true, nil
}

func QueryFavoriteVideoByUserId(ctx context.Context, request *douyinfavorite.DouyinFavoriteListRequest) (*[]int64, error) {
	var fav *[]Favorite_db
	result := DB.WithContext(ctx).Where("user_id = ?", request.UserId).Order("favorite_time desc").Find(&fav)
	if result.Error != nil {
		return nil, result.Error
	}

	var videoIds []int64
	for _, v := range *fav {

		videoIds = append(videoIds, v.VideoId)
	}
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
