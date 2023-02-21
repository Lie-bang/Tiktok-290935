package main

import (
	"context"
	"douyin/cmd/video/Service"
	douyinvideo "douyin/kitex_gen/douyinvideo"
	"fmt"
	"log"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// FeedVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FeedVideo(ctx context.Context, request *douyinvideo.DouyinFeedRequest) (resp *douyinvideo.DouyinFeedResponse, err error) {
	// TODO: Your code here...
	v, respTime, err := Service.NewFeedVideoService(ctx).FeedVideo(request)
	resp = new(douyinvideo.DouyinFeedResponse)
	if err != nil {
		fmt.Println(err)
		resp.StatusCode = 1
		return resp, nil
	}

	if v == nil {
		resp.StatusCode = 0
		resp.VideoList = v
		var statusMsg string
		statusMsg = "No Video Yet, Please Upload your video First."
		resp.NextTime = &respTime
		resp.StatusMsg = &statusMsg
		return resp, nil
	}

	resp.StatusCode = 0
	resp.VideoList = v
	var statusMsg string
	statusMsg = "success"
	resp.NextTime = &respTime
	resp.StatusMsg = &statusMsg

	return resp, nil
}

// PublishVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishVideo(ctx context.Context, request *douyinvideo.DouyinPublishActionRequest) (resp *douyinvideo.DouyinPublishActionResponse, err error) {
	// TODO: Your code here...
	err = Service.NewPublishVideoService(ctx).PublishVideo(request)
	if err != nil {
		log.Print(err)
		return nil, nil
	}
	resp = new(douyinvideo.DouyinPublishActionResponse)
	statusCode := 0
	resp.StatusCode = int32(statusCode)
	msg := "success"
	resp.StatusMsg = &msg

	return resp, nil
}

// PublishListVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishListVideo(ctx context.Context, request *douyinvideo.DouyinPublishListRequest) (resp *douyinvideo.DouyinPublishListResponse, err error) {
	// TODO: Your code here...
	v, err := Service.NewPublishListVideoService(ctx).PublishListVideo(request)
	resp = new(douyinvideo.DouyinPublishListResponse)
	if err != nil {
		log.Print(err)
		return resp, err
	}
	resp.StatusCode = 0
	statusMsg := "Get video list success"
	resp.StatusMsg = &statusMsg
	resp.VideoList = v

	return resp, err
}

// FavoriteVideoList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FavoriteVideoList(ctx context.Context, request *douyinvideo.DouyinFavoriteListRequest) (resp *douyinvideo.DouyinFavoriteListResponse, err error) {
	// TODO: Your code here...
	v, err := Service.NewFavoriteVideoListService(ctx).FavoriteVideoList(request)

	if err != nil {
		log.Print(err)
		return nil, err
	}
	resp = new(douyinvideo.DouyinFavoriteListResponse)
	resp.StatusCode = 0
	statusMsg := "Get video list success"
	resp.StatusMsg = &statusMsg
	resp.VideoList = v
	return resp, err
}

// CommentCountUpdate implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CommentCountUpdate(ctx context.Context, request *douyinvideo.DouyinCommentCountRequest) (resp *douyinvideo.DouyinCommentCountResponse, err error) {
	// TODO: Your code here...
	err = Service.NewUpdateCommentCountService(ctx).UpdateCommentCount(request)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	resp = new(douyinvideo.DouyinCommentCountResponse)
	resp.StatusCode = 0
	StatusMsg := "update comment count success"
	resp.StatusMsg = &StatusMsg

	return resp, nil
}

// WorkAndFavoriteCount implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) WorkAndFavoriteCount(ctx context.Context, request *douyinvideo.Douyin_Work_And_Favorite_CountRequest) (resp *douyinvideo.Douyin_Work_And_Favorite_CountResponse, err error) {
	// TODO: Your code here...
	workCount, favCount, err := Service.NewWorkAndFavoriteCountService(ctx).WorkAndFavoriteCount(request)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	resp = new(douyinvideo.Douyin_Work_And_Favorite_CountResponse)
	resp.WorkCount = workCount
	resp.FavCount = favCount
	resp.StatusCode = 0
	StatusMsg := "update comment count success"
	resp.StatusMsg = &StatusMsg

	return resp, err
}
