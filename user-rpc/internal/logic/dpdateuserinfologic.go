package logic

import (
	"context"

	"go-zero-looklook/user-rpc/internal/svc"
	"go-zero-looklook/user-rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DpdateUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DpdateUserInfoLogic {
	return &DpdateUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DpdateUserInfoLogic) DpdateUserInfo(in *pb.UpdateUserInfoReq) (*pb.UpdateUserInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateUserInfoResp{}, nil
}
