namespace go snapshot

struct BaseReq{
    1: string LogId = ""
}

struct BaseResp{
    1: string StatusMessage = ""
    2: i32 StatusCode
}

struct PreCreateOrderReq {
    1: BaseReq req;
    2: string orderInfo;
    3: string conId;
    4: string userId;
}

struct PreCreateOrderResp {
    1: BaseResp resp;
}

struct CoreCreateOrderReq {
    1: BaseReq Base;
    2: i64 tokenReq;

}

struct CoreCreateOrderResp {
    1: BaseResp resp;
}

service OrderServiceFacade {
    PreCreateOrderResp PreCreateOrder(1: PreCreateOrderReq req);
    CoreCreateOrderResp CoreCreateOrder(1: CoreCreateOrderReq req);
}
