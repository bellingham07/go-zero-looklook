package logic

import (
	"context"
	"fmt"

	"go-zero-looklook/user-rpc/internal/svc"
	"go-zero-looklook/user-rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	m := map[int64]string{
		1: "zhangan",
		2: "lisi",
	}

	fmt.Println(in.Id, "--------------------")

	nickname := "unknown"
	if name, ok := m[in.Id]; ok {
		nickname = name
	}
	return &pb.GetUserInfoResp{
		Id:       in.Id,
		Nickname: nickname,
	}, nil
}
