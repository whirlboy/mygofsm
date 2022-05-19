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
	ExtraInfo   map[interface{}]interface{}
}

type OrderPayment struct {
	OrderPaymentId int64
	OrderId        int64
	PaymentStatus  int64
	PaymentAmount  int64
	ExtraInfo      map[interface{}]interface{}
}

type OrderBrief struct {
	OrderId     int64
	OrderTime   string
	ConId       int64
	OrderDetail string
	ExtraInfo   map[interface{}]interface{}
}
