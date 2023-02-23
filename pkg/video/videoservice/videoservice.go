// Code generated by goctl. DO NOT EDIT.
// Source: video.proto

package videoservice

import (
	"context"

	"douyin/pkg/video/types/video"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ChangeVideoCommentReq   = video.ChangeVideoCommentReq
	ChangeVideoCommentResp  = video.ChangeVideoCommentResp
	ChangeVideoFavoriteReq  = video.ChangeVideoFavoriteReq
	ChangeVideoFavoriteResp = video.ChangeVideoFavoriteResp
	GetAllVideoByUserIdReq  = video.GetAllVideoByUserIdReq
	GetAllVideoByUserIdResp = video.GetAllVideoByUserIdResp
	GetVideoReq             = video.GetVideoReq
	GetVideoResp            = video.GetVideoResp
	PublishVideoReq         = video.PublishVideoReq
	PublishVideoResp        = video.PublishVideoResp
	User                    = video.User
	Video                   = video.Video

	VideoService interface {
		PublishVideo(ctx context.Context, in *PublishVideoReq, opts ...grpc.CallOption) (*PublishVideoResp, error)
		GetVideo(ctx context.Context, in *GetVideoReq, opts ...grpc.CallOption) (*GetVideoResp, error)
		GetAllVideoByUserId(ctx context.Context, in *GetAllVideoByUserIdReq, opts ...grpc.CallOption) (*GetAllVideoByUserIdResp, error)
		ChangeVideoComment(ctx context.Context, in *ChangeVideoCommentReq, opts ...grpc.CallOption) (*ChangeVideoCommentResp, error)
		ChangeVideoFavorite(ctx context.Context, in *ChangeVideoFavoriteReq, opts ...grpc.CallOption) (*ChangeVideoFavoriteResp, error)
	}

	defaultVideoService struct {
		cli zrpc.Client
	}
)

func NewVideoService(cli zrpc.Client) VideoService {
	return &defaultVideoService{
		cli: cli,
	}
}

func (m *defaultVideoService) PublishVideo(ctx context.Context, in *PublishVideoReq, opts ...grpc.CallOption) (*PublishVideoResp, error) {
	client := video.NewVideoServiceClient(m.cli.Conn())
	return client.PublishVideo(ctx, in, opts...)
}

func (m *defaultVideoService) GetVideo(ctx context.Context, in *GetVideoReq, opts ...grpc.CallOption) (*GetVideoResp, error) {
	client := video.NewVideoServiceClient(m.cli.Conn())
	return client.GetVideo(ctx, in, opts...)
}

func (m *defaultVideoService) GetAllVideoByUserId(ctx context.Context, in *GetAllVideoByUserIdReq, opts ...grpc.CallOption) (*GetAllVideoByUserIdResp, error) {
	client := video.NewVideoServiceClient(m.cli.Conn())
	return client.GetAllVideoByUserId(ctx, in, opts...)
}

func (m *defaultVideoService) ChangeVideoComment(ctx context.Context, in *ChangeVideoCommentReq, opts ...grpc.CallOption) (*ChangeVideoCommentResp, error) {
	client := video.NewVideoServiceClient(m.cli.Conn())
	return client.ChangeVideoComment(ctx, in, opts...)
}

func (m *defaultVideoService) ChangeVideoFavorite(ctx context.Context, in *ChangeVideoFavoriteReq, opts ...grpc.CallOption) (*ChangeVideoFavoriteResp, error) {
	client := video.NewVideoServiceClient(m.cli.Conn())
	return client.ChangeVideoFavorite(ctx, in, opts...)
}
