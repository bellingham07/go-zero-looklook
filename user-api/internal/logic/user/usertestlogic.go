package user

import (
	"context"
	"fmt"
	"go-zero-looklook/user-api/internal/model"
	"go-zero-looklook/user-api/internal/svc"
	"go-zero-looklook/user-api/internal/types"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserTestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserTestLogic {
	return &UserTestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserTestLogic) UserTest(req *types.LoginReq) (resp *types.LoginResp, err error) {
	var t model.Test
	l.svcCtx.DB.Debug().Where("id = ?", 1).Find(&t)
	var tt int64
	if t.Time.Valid {
		tt = t.Time.Int64
	} else {
		tt = 0
	}
	fmt.Println(time.Unix(tt, 0).Format("2006:01:02 15:04:05"))
	return
}
