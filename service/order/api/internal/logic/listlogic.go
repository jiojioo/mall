package logic

import (
	"context"
	"mall/service/order/rpc/types/order"

	"mall/service/order/api/internal/svc"
	"mall/service/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.ListRequest) (resp *types.ListResponse, err error) {
	res, err := l.svcCtx.OrderRpc.List(l.ctx, &order.ListRequest{
		Uid: int64(req.Uid),
	})
	if err != nil {
		return nil, err
	}
	orderList := make([]*types.DetailResponse, 0)
	for _, item := range res.Data {
		orderList = append(orderList, &types.DetailResponse{
			Id:     item.Id,
			Uid:    item.Uid,
			Pid:    item.Pid,
			Amount: item.Amount,
			Status: item.Status,
		})
	}
	return &types.ListResponse{
		Data:  orderList,
		Total: int64(len(orderList)),
	}, nil
}
