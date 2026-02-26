namespace go order
include "../base/base.thrift"
include "../base/order.thrift"

struct GetOrderReq {
    1:optional i64 id=0 (api.raw = "id")
    2:optional string sn="" (api.raw = "sn")
}
struct GetOrderListReq{
    1:optional i64 memberId=0 (api.raw = "memberId")
    2:optional i64 createdId=0 (api.raw = "createdId")
    3:optional string status = "" (api.raw = "status")
    4:optional string nature="" (api.raw = "nature")
    5:optional string searchKey="" (api.raw = "searchKey")

    254: optional i64 page=1
    255: optional i64 pageSize=100
}

struct RefundOrderReq {
    1:optional i64 id=0 (api.raw = "id")
    2:optional string reason="" (api.raw = "reason")
    3:optional i64 amount=0 (api.raw = "amount")
    4:optional i64 createdId=0 (api.raw = "createdId")
}

struct CreateOrderReq {
    1:optional i64 memberId=0 (api.raw = "memberId")
    2:optional i64 createdId=0 (api.raw = "createdId")
    3:optional list<order.OrderItem> items=[] (api.raw = "items")
    4:optional i64 totalAmount=0 (api.raw = "totalAmount")
   5:optional i64 userId=0 (api.raw = "userId")
}
struct CancelledOrderReq {
    1:optional i64 id=0 (api.raw = "id")
    2:optional string reason="" (api.raw = "reason")
    3:optional i64 createdId=0 (api.raw = "createdId")
}
struct PaymentReq {

}
service OrderService {
    base.NilResponse GetOrder(1:GetOrderReq req)(api.post = "/service/order")
    base.NilResponse GetOrderList(1:GetOrderListReq req)(api.post = "/service/order/list")
    base.NilResponse DeleteOrder(1:base.IdReq req)(api.post = "/service/order/delete")
    base.NilResponse CreateOrder(1:CreateOrderReq req)(api.post = "/service/order/create")
    base.NilResponse Payment(1:PaymentReq req)(api.post = "/service/order/payment")
    base.NilResponse CancelledOrder(1:CancelledOrderReq req)(api.post = "/service/order/cancelled")
    base.NilResponse RefundOrder(1:RefundOrderReq req)(api.post = "/service/order/refund")

 }



