// Code generated by Kitex v0.4.4. DO NOT EDIT.

package videoservice

import (
	"context"
	douyinvideo "douyin/kitex_gen/douyinvideo"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return videoServiceServiceInfo
}

var videoServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "VideoService"
	handlerType := (*douyinvideo.VideoService)(nil)
	methods := map[string]kitex.MethodInfo{
		"FeedVideo":            kitex.NewMethodInfo(feedVideoHandler, newVideoServiceFeedVideoArgs, newVideoServiceFeedVideoResult, false),
		"PublishVideo":         kitex.NewMethodInfo(publishVideoHandler, newVideoServicePublishVideoArgs, newVideoServicePublishVideoResult, false),
		"PublishListVideo":     kitex.NewMethodInfo(publishListVideoHandler, newVideoServicePublishListVideoArgs, newVideoServicePublishListVideoResult, false),
		"FavoriteVideoList":    kitex.NewMethodInfo(favoriteVideoListHandler, newVideoServiceFavoriteVideoListArgs, newVideoServiceFavoriteVideoListResult, false),
		"CommentCountUpdate":   kitex.NewMethodInfo(commentCountUpdateHandler, newVideoServiceCommentCountUpdateArgs, newVideoServiceCommentCountUpdateResult, false),
		"WorkAndFavoriteCount": kitex.NewMethodInfo(workAndFavoriteCountHandler, newVideoServiceWorkAndFavoriteCountArgs, newVideoServiceWorkAndFavoriteCountResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "douyinvideo",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func feedVideoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*douyinvideo.VideoServiceFeedVideoArgs)
	realResult := result.(*douyinvideo.VideoServiceFeedVideoResult)
	success, err := handler.(douyinvideo.VideoService).FeedVideo(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceFeedVideoArgs() interface{} {
	return douyinvideo.NewVideoServiceFeedVideoArgs()
}

func newVideoServiceFeedVideoResult() interface{} {
	return douyinvideo.NewVideoServiceFeedVideoResult()
}

func publishVideoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*douyinvideo.VideoServicePublishVideoArgs)
	realResult := result.(*douyinvideo.VideoServicePublishVideoResult)
	success, err := handler.(douyinvideo.VideoService).PublishVideo(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServicePublishVideoArgs() interface{} {
	return douyinvideo.NewVideoServicePublishVideoArgs()
}

func newVideoServicePublishVideoResult() interface{} {
	return douyinvideo.NewVideoServicePublishVideoResult()
}

func publishListVideoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*douyinvideo.VideoServicePublishListVideoArgs)
	realResult := result.(*douyinvideo.VideoServicePublishListVideoResult)
	success, err := handler.(douyinvideo.VideoService).PublishListVideo(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServicePublishListVideoArgs() interface{} {
	return douyinvideo.NewVideoServicePublishListVideoArgs()
}

func newVideoServicePublishListVideoResult() interface{} {
	return douyinvideo.NewVideoServicePublishListVideoResult()
}

func favoriteVideoListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*douyinvideo.VideoServiceFavoriteVideoListArgs)
	realResult := result.(*douyinvideo.VideoServiceFavoriteVideoListResult)
	success, err := handler.(douyinvideo.VideoService).FavoriteVideoList(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceFavoriteVideoListArgs() interface{} {
	return douyinvideo.NewVideoServiceFavoriteVideoListArgs()
}

func newVideoServiceFavoriteVideoListResult() interface{} {
	return douyinvideo.NewVideoServiceFavoriteVideoListResult()
}

func commentCountUpdateHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*douyinvideo.VideoServiceCommentCountUpdateArgs)
	realResult := result.(*douyinvideo.VideoServiceCommentCountUpdateResult)
	success, err := handler.(douyinvideo.VideoService).CommentCountUpdate(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceCommentCountUpdateArgs() interface{} {
	return douyinvideo.NewVideoServiceCommentCountUpdateArgs()
}

func newVideoServiceCommentCountUpdateResult() interface{} {
	return douyinvideo.NewVideoServiceCommentCountUpdateResult()
}

func workAndFavoriteCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*douyinvideo.VideoServiceWorkAndFavoriteCountArgs)
	realResult := result.(*douyinvideo.VideoServiceWorkAndFavoriteCountResult)
	success, err := handler.(douyinvideo.VideoService).WorkAndFavoriteCount(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceWorkAndFavoriteCountArgs() interface{} {
	return douyinvideo.NewVideoServiceWorkAndFavoriteCountArgs()
}

func newVideoServiceWorkAndFavoriteCountResult() interface{} {
	return douyinvideo.NewVideoServiceWorkAndFavoriteCountResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) FeedVideo(ctx context.Context, request *douyinvideo.DouyinFeedRequest) (r *douyinvideo.DouyinFeedResponse, err error) {
	var _args douyinvideo.VideoServiceFeedVideoArgs
	_args.Request = request
	var _result douyinvideo.VideoServiceFeedVideoResult
	if err = p.c.Call(ctx, "FeedVideo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PublishVideo(ctx context.Context, request *douyinvideo.DouyinPublishActionRequest) (r *douyinvideo.DouyinPublishActionResponse, err error) {
	var _args douyinvideo.VideoServicePublishVideoArgs
	_args.Request = request
	var _result douyinvideo.VideoServicePublishVideoResult
	if err = p.c.Call(ctx, "PublishVideo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PublishListVideo(ctx context.Context, request *douyinvideo.DouyinPublishListRequest) (r *douyinvideo.DouyinPublishListResponse, err error) {
	var _args douyinvideo.VideoServicePublishListVideoArgs
	_args.Request = request
	var _result douyinvideo.VideoServicePublishListVideoResult
	if err = p.c.Call(ctx, "PublishListVideo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FavoriteVideoList(ctx context.Context, request *douyinvideo.DouyinFavoriteListRequest) (r *douyinvideo.DouyinFavoriteListResponse, err error) {
	var _args douyinvideo.VideoServiceFavoriteVideoListArgs
	_args.Request = request
	var _result douyinvideo.VideoServiceFavoriteVideoListResult
	if err = p.c.Call(ctx, "FavoriteVideoList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CommentCountUpdate(ctx context.Context, request *douyinvideo.DouyinCommentCountRequest) (r *douyinvideo.DouyinCommentCountResponse, err error) {
	var _args douyinvideo.VideoServiceCommentCountUpdateArgs
	_args.Request = request
	var _result douyinvideo.VideoServiceCommentCountUpdateResult
	if err = p.c.Call(ctx, "CommentCountUpdate", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) WorkAndFavoriteCount(ctx context.Context, request *douyinvideo.Douyin_Work_And_Favorite_CountRequest) (r *douyinvideo.Douyin_Work_And_Favorite_CountResponse, err error) {
	var _args douyinvideo.VideoServiceWorkAndFavoriteCountArgs
	_args.Request = request
	var _result douyinvideo.VideoServiceWorkAndFavoriteCountResult
	if err = p.c.Call(ctx, "WorkAndFavoriteCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}