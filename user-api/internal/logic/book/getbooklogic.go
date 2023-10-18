package book

import (
	"context"
	"database/sql"
	"go-zero-looklook/user-api/internal/model"
	"time"

	"go-zero-looklook/user-api/internal/svc"
	"go-zero-looklook/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBookLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBookLogic {
	return &GetBookLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBookLogic) GetBook(req *types.BookReq) (resp *types.BookRepl, err error) {
	// 将时间戳转为时间
	t := &model.Test{
		Id:   1,
		Name: sql.NullString{req.Name, true},
		Time: sql.NullInt64{time.Now().Unix(), true},
	}
	// 存储到数据库
	l.svcCtx.DB.Debug().Save(&t)
	return
}
