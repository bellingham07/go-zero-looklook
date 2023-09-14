package user

import (
	"context"
	"github.com/pkg/errors"
	"go-zero-looklook/user-api/internal/svc"
	"go-zero-looklook/user-api/internal/types"

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
	//logx.Error("nihaoa")
	//resp.Id = 1
	//if err = l.test1(); err != nil {
	//	logx.Errorf("%+v \n", err)
	//}
	resp.Id = 0
	return
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
