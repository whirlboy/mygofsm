package repository

import (
	"context"
	"mygofsm/my_order_project/domain/order_domain/order_aggregate/entity"
)

/**
 * @Author waizixi
 * @Description //TODO $
 **/

var OrderRepo OrderRepository

type OrderRepository interface {
	CreateOrder(ctx *context.Context, orderEntity *entity.Order) error
	GetOrderEntityById(ctx *context.Context, orderId int64) (*entity.Order, error)
	UpdateOrderStatus(ctx *context.Context) error
}
