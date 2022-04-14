package order_aggregate

import (
	"context"
	"mygofsm/my_order_project/domain/order_domain/order_aggregate/entity"
)

/**
 * @Author waizixi
 * @Description //TODO $
 **/

type OrderRepositoryImpl struct {
}

func (impl *OrderRepositoryImpl) CreateOrder(ctx *context.Context, orderEntity *entity.Order) error {
	return nil
}
func (impl *OrderRepositoryImpl) GetOrderEntityById(ctx *context.Context, orderId int64) (*entity.Order, error) {
	return nil, nil
}
func (impl *OrderRepositoryImpl) UpdateOrderStatus(ctx *context.Context) error {
	return nil
}
