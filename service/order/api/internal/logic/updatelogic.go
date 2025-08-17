package logic

import (
	"context"
	"mall/service/order/rpc/types/order"

	"mall/service/order/api/internal/svc"
	"mall/service/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UpdateRequest) (resp *types.UpdateResponse, err error) {
	res, err := l.svcCtx.OrderRpc.Update(l.ctx, &order.UpdateRequest{
		Id:     req.Id,
		Uid:    req.Uid,
		Pid:    req.Pid,
		Amount: req.Amount,
		Status: req.Status,
	})

	if err != nil {
		return nil, err
	}

	return &types.UpdateResponse{
		Id:     res.Id,
		Uid:    res.Uid,
		Pid:    res.Pid,
		Amount: res.Amount,
		Status: res.Status,
	}, nil
	//var orderReq order.UpdateRequest
	//if err := copier.Copy(&orderReq, req); err != nil {
	//	return nil, fmt.Errorf("复制请求参数失败: %w", err) // 包装错误，增加上下文
	//}
	//
	//// 调用 Rpc 服务
	//res, err := l.svcCtx.OrderRpc.Update(l.ctx, &orderReq)
	//if err != nil {
	//	return nil, fmt.Errorf("调用 OrderRpc.Update 失败: %w", err) // 包装错误
	//}
	//
	//// 自动复制 res 到 types.UpdateResponse
	//var updateResp types.UpdateResponse
	//if err := copier.Copy(&updateResp, res); err != nil {
	//	return nil, fmt.Errorf("复制响应结果失败: %w", err) // 包装错误
	//}
	//
	//return &updateResp, nil
}
