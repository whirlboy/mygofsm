package entity

/**
 * @Author waizixi
 * @Description //TODO $
 **/

type Order struct {
	OrderId     int64 //订单id
	OrderType   int64
	OrderName   string
	OrderStatus int64
	AuditStatus int64
	SkuIds      []int64
	PayMode     int64
	TaskIds     []int64
}