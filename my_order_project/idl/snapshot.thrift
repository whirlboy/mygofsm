namespace go snapshot

struct PreSubmitOrdersReq {
    1: string msg;
}

struct PreSubmitOrdersResp {
    1: string msg;
}

service OrderService {
    PreSubmitOrdersResp PreSubmitOrders(1: PreSubmitOrdersReq req);
}
