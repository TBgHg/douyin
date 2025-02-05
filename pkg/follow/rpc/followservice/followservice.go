// Code generated by goctl. DO NOT EDIT.
// Source: follow.proto

package followservice

import (
	"context"

	"douyin/pkg/follow/rpc/types/follow"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	FollowReq           = follow.FollowReq
	FollowResp          = follow.FollowResp
	GetFollowListReq    = follow.GetFollowListReq
	GetFollowListResp   = follow.GetFollowListResp
	GetFollowerListReq  = follow.GetFollowerListReq
	GetFollowerListResp = follow.GetFollowerListResp
	User                = follow.User

	FollowService interface {
		Follow(ctx context.Context, in *FollowReq, opts ...grpc.CallOption) (*FollowResp, error)
		GetFollowList(ctx context.Context, in *GetFollowListReq, opts ...grpc.CallOption) (*GetFollowListResp, error)
		GetFollowerList(ctx context.Context, in *GetFollowerListReq, opts ...grpc.CallOption) (*GetFollowerListResp, error)
	}

	defaultFollowService struct {
		cli zrpc.Client
	}
)

func NewFollowService(cli zrpc.Client) FollowService {
	return &defaultFollowService{
		cli: cli,
	}
}

func (m *defaultFollowService) Follow(ctx context.Context, in *FollowReq, opts ...grpc.CallOption) (*FollowResp, error) {
	client := follow.NewFollowServiceClient(m.cli.Conn())
	return client.Follow(ctx, in, opts...)
}

func (m *defaultFollowService) GetFollowList(ctx context.Context, in *GetFollowListReq, opts ...grpc.CallOption) (*GetFollowListResp, error) {
	client := follow.NewFollowServiceClient(m.cli.Conn())
	return client.GetFollowList(ctx, in, opts...)
}

func (m *defaultFollowService) GetFollowerList(ctx context.Context, in *GetFollowerListReq, opts ...grpc.CallOption) (*GetFollowerListResp, error) {
	client := follow.NewFollowServiceClient(m.cli.Conn())
	return client.GetFollowerList(ctx, in, opts...)
}
