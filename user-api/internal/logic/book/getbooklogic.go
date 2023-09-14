package book

import (
	"context"

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
	// todo: add your logic here and delete this line
	resp.Id = "0"
	return
}
