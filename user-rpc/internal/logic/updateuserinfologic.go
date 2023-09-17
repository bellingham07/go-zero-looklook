package logic

import (
	"context"
	"strconv"

	"go-zero-looklook/user-rpc/internal/svc"
	"go-zero-looklook/user-rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(in *pb.UpdateUserInfoReq) (*pb.UpdateUserInfoResp, error) {
	name := strconv.Itoa(int(in.Id)) + "caojinbo"
	return &pb.UpdateUserInfoResp{Info: name}, nil
}
