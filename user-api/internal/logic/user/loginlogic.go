package user

import (
	"context"
	"github.com/pkg/errors"
	"go-zero-looklook/user-api/internal/svc"
	"go-zero-looklook/user-api/internal/types"
	"go-zero-looklook/user-rpc/pb/pb"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	userRpc, err := l.svcCtx.UserRpcClient.GetUserInfo(l.ctx, &pb.GetUserInfoReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &types.LoginResp{
		Id:   userRpc.Id,
		Name: userRpc.Nickname,
		Time: time.Now().String(),
	}, nil
}

func (l *LoginLogic) test1() error {
	return l.test2()
}
func (l *LoginLogic) test2() error {
	return l.test3()
}
func (l *LoginLogic) test3() error {
	return errors.Wrap(errors.New("这是故意的"), "hahah")
}
