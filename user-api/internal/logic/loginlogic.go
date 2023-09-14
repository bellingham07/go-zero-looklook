package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go-zero-looklook/user-api/internal/model"
	"go-zero-looklook/user-api/internal/svc"
	"go-zero-looklook/user-api/internal/types"
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
	fmt.Println("12321424-----------------------------------")
	var user model.Test
	if err = l.svcCtx.DB.Debug().Table("test").Find(&user).Error; err != nil {
		err = errors.New("shujuk")
		return
	}
	fmt.Printf("%+v", user)
	var t string
	var a int64
	if user.Time.Valid {
		t = user.Time.Time.Format("2006-01-02 15:04:05")
		a = user.Time.Time.Unix()
		fmt.Println(a)
	}
	b := time.Unix(a, 0)
	fmt.Println(b)
	newUser := &model.Test{
		Id:   2,
		Name: sql.NullString{"123", true},
		Time: sql.NullTime{time.Now(), true},
	}
	l.svcCtx.DB.Save(&newUser)

	return &types.LoginResp{
		Id:   user.Id,
		Name: user.Name.String,
		Time: t,
	}, err
}
