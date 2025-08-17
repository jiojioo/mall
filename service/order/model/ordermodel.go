package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OrderModel = (*customOrderModel)(nil)

type (
	// OrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderModel.
	OrderModel interface {
		orderModel
		FindAllByUid(ctx context.Context, uid int64) ([]*Order, error)
	}

	customOrderModel struct {
		*defaultOrderModel
	}
)

// NewOrderModel returns a model for the database table.
func NewOrderModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) OrderModel {
	return &customOrderModel{
		defaultOrderModel: newOrderModel(conn, c, opts...),
	}
}

func (m *customOrderModel) FindAllByUid(ctx context.Context, uid int64) ([]*Order, error) {
	var resp []*Order

	query := fmt.Sprintf("select %s from %s where `uid` = ?", orderRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, uid)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

//func (m *customOrderModel) FindAllByUid(ctx context.Context, uid int64, page, size int) ([]*Order, int64, error) {
//	// 1. 查询总条数（用于计算总页数）
//	var total int64
//	countQuery := fmt.Sprintf("select count(*) from %s where uid = ?", m.table)
//	if err := m.QueryRowCtx(ctx, &total, countQuery, uid); err != nil {
//		return nil, 0, err
//	}
//
//	// 2. 计算偏移量（分页公式：offset = (page-1)*size）
//	offset := (page - 1) * size
//	query := fmt.Sprintf("select %s from %s where uid = ? limit ? offset ?", orderRows, m.table)
//
//	// 3. 分页查询
//	var resp []*Order
//	if err := m.QueryRowsNoCacheCtx(ctx, &resp, query, uid, size, offset); err != nil {
//		return nil, 0, err
//	}
//
//	return resp, total, nil
//}
