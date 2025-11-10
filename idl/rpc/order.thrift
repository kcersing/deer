namespace go order
include "../base/base.thrift"
include "../base/order.thrift"





struct GetOrderInfoReq {
    1:optional i64 id=0 (api.raw = "id")
    2:optional string sn (api.raw = "sn")
}
struct GetOrderListReq{
    1:optional i64 memberId=0 (api.raw = "memberId")
    2:optional i64 createdId=0 (api.raw = "createdId")
    3:optional string status(api.raw = "status")
    4:optional string nature="" (api.raw = "nature")
    5:optional string searchKey="" (api.raw = "searchKey")

    254:optional base.PageReq pageReq = {}
}
struct GetOrderListResp {
    1:optional list<order.Order> data={}
    255:optional base.BaseResp baseResp={}
}

struct RefundOrderReq {
    1:optional i64 id=0 (api.raw = "id")
    2:optional string reason="" (api.raw = "reason")
    3:optional i64 amount=0 (api.raw = "amount")
    4:optional i64 createdId=0 (api.raw = "createdId")
}

struct CancelledOrderReq {
    1:optional i64 id=0 (api.raw = "id")
    2:optional string reason="" (api.raw = "reason")
    3:optional i64 createdId=0 (api.raw = "createdId")
}

struct CreateOrderResp {
    1:optional order.Order data={}
    255:optional base.BaseResp baseResp={}
}

struct OrderResp {
    1:optional order.Order data={}
    255:optional base.BaseResp baseResp={}
}
struct CreateOrderReq {
    1:optional i64 memberId=0 (api.raw = "memberId")
    2:optional i64 createdId=0 (api.raw = "createdId")
    3:optional order.OrderItem items={} (api.raw = "items")
    4:optional i64 totalAmount=0 (api.raw = "totalAmount")
}

struct PaymentReq {

}

service OrderService  {
    OrderResp GetOrderInfo(1:GetOrderInfoReq req)
    GetOrderListResp GetOrderList(1:GetOrderListReq req)
    base.NilResponse DeleteOrder(1:base.IdReq req)
    OrderResp CreateOrder(1:GetOrderListReq req)
    OrderResp Payment(1:PaymentReq req)
    base.NilResponse CancelledOrder(1:CreateOrderReq req)
    base.NilResponse RefundOrder(1:RefundOrderReq req)
}