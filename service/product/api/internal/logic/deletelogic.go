package logic

import (
	"context"
	"mall/service/product/rpc/types/product"

	"mall/service/product/api/internal/svc"
	"mall/service/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.DeleteResquest) (resp *types.DeleteResponse, err error) {
	res, err := l.svcCtx.ProductRpc.Delete(l.ctx, &product.DeleteRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.DeleteResponse{
		Result: res.Result,
	}, nil
}
