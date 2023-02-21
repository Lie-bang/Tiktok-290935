// Code generated by Kitex v0.4.4. DO NOT EDIT.

package userservice

import (
	"context"
	douyinuser "douyin/kitex_gen/douyinuser"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateUser(ctx context.Context, req *douyinuser.CreateUserRequest, callOptions ...callopt.Option) (r *douyinuser.CreateUserResponse, err error)
	CheckUser(ctx context.Context, req *douyinuser.CheckUserRequest, callOptions ...callopt.Option) (r *douyinuser.CheckUserResponse, err error)
	GetUser(ctx context.Context, req *douyinuser.GetUserRequest, callOptions ...callopt.Option) (r *douyinuser.GetUserResponse, err error)
	MGetUserName(ctx context.Context, req *douyinuser.MGetUserNameRequest, callOptions ...callopt.Option) (r *douyinuser.MGetUserNameResponse, err error)
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
	return &kUserServiceClient{
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

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) CreateUser(ctx context.Context, req *douyinuser.CreateUserRequest, callOptions ...callopt.Option) (r *douyinuser.CreateUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateUser(ctx, req)
}

func (p *kUserServiceClient) CheckUser(ctx context.Context, req *douyinuser.CheckUserRequest, callOptions ...callopt.Option) (r *douyinuser.CheckUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CheckUser(ctx, req)
}

func (p *kUserServiceClient) GetUser(ctx context.Context, req *douyinuser.GetUserRequest, callOptions ...callopt.Option) (r *douyinuser.GetUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUser(ctx, req)
}

func (p *kUserServiceClient) MGetUserName(ctx context.Context, req *douyinuser.MGetUserNameRequest, callOptions ...callopt.Option) (r *douyinuser.MGetUserNameResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MGetUserName(ctx, req)
}
