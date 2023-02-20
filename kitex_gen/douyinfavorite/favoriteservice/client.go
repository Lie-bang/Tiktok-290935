// Code generated by Kitex v0.4.4. DO NOT EDIT.

package favoriteservice

import (
	"context"
	douyinfavorite "douyin/kitex_gen/douyinfavorite"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	FavoriteAction(ctx context.Context, request *douyinfavorite.DouyinFavoriteActionRequest, callOptions ...callopt.Option) (r *douyinfavorite.DouyinFavoriteActionResponse, err error)
	FavoriteList(ctx context.Context, request *douyinfavorite.DouyinFavoriteListRequest, callOptions ...callopt.Option) (r *douyinfavorite.DouyinFavoriteListResponse, err error)
	FavoriteJudge(ctx context.Context, request *douyinfavorite.DouyinFavoriteJudgeRequest, callOptions ...callopt.Option) (r *douyinfavorite.DouyinFavoriteJudgeResponse, err error)
	FavoriteCountUser(ctx context.Context, request *douyinfavorite.DouyinFavoriteCountUserRequest, callOptions ...callopt.Option) (r *douyinfavorite.DouyinFavoriteCountUserResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kFavoriteServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kFavoriteServiceClient struct {
	*kClient
}

func (p *kFavoriteServiceClient) FavoriteAction(ctx context.Context, request *douyinfavorite.DouyinFavoriteActionRequest, callOptions ...callopt.Option) (r *douyinfavorite.DouyinFavoriteActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteAction(ctx, request)
}

func (p *kFavoriteServiceClient) FavoriteList(ctx context.Context, request *douyinfavorite.DouyinFavoriteListRequest, callOptions ...callopt.Option) (r *douyinfavorite.DouyinFavoriteListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteList(ctx, request)
}

func (p *kFavoriteServiceClient) FavoriteJudge(ctx context.Context, request *douyinfavorite.DouyinFavoriteJudgeRequest, callOptions ...callopt.Option) (r *douyinfavorite.DouyinFavoriteJudgeResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteJudge(ctx, request)
}

func (p *kFavoriteServiceClient) FavoriteCountUser(ctx context.Context, request *douyinfavorite.DouyinFavoriteCountUserRequest, callOptions ...callopt.Option) (r *douyinfavorite.DouyinFavoriteCountUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteCountUser(ctx, request)
}