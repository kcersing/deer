namespace go deer.order
include "base.thrift"


struct Item {
    1:optional i64 productId=0 (api.raw = "productId")
    2:optional i64 quantity=1 (api.raw = "quantity")
    3:optional double price=0 (api.raw = "price")
    4:optional string name="" (api.raw = "name")
}
struct Order {
    1:optional i64 memberId=0 (api.raw = "memberId")
    2:optional i64 createdId=1 (api.raw = "createdId")

    3:optional list<Item> items={} (api.raw = "items")
    4:optional string sn="" (api.raw = "sn")
    5:optional double totalAmount=0 (api.raw = "totalAmount")
    6:optional string status="" (api.raw = "status")
    7:optional string nature="" (api.raw = "nature")

    8:optional string createdAt="" (api.raw = "createdAt")
    9:optional string completionAt="" (api.raw = "completionAt")
    10:optional string closeAt="" (api.raw = "closeAt")
    11:optional string refundAt="" (api.raw = "refundAt")
//    12:optional i64 version=0 (api.raw = "version")
    13:optional string updatedAt="" (api.raw = "updatedAt")

    254:optional i64 id=0 (api.raw = "id")
}


struct GetOrderInfoReq {
    1:optional i64 id=0 (api.raw = "id")
    2:optional string sn (api.raw = "sn")
}
struct GetOrderListReq{
    1:optional i64 memberId=0 (api.raw = "memberId")
    2:optional i64 createdId=1 (api.raw = "createdId")
    3:optional string status(api.raw = "status")
    4:optional string nature="" (api.raw = "nature")
    5:optional string searchKey="" (api.raw = "searchKey")

    254:optional base.PageReq pageReq
}
struct GetOrderListResp {
    1:optional list<Order> data
    255:optional base.BaseResp baseResp
}

struct RefundOrderReq {
    1:optional i64 id=0 (api.raw = "id")
    2:optional string reason="" (api.raw = "reason")
    3:optional double amount=0 (api.raw = "amount")
    4:optional i64 createdId=1 (api.raw = "createdId")
}

struct CancelledOrderReq {
    1:optional i64 id=0 (api.raw = "id")
    2:optional string reason="" (api.raw = "reason")
    3:optional i64 createdId=1 (api.raw = "createdId")
}

struct CreateOrderResp {
    1:optional Order order
    255:optional base.BaseResp baseResp
}

struct OrderResp {
    1:optional Order order
    255:optional base.BaseResp baseResp
}
struct CreateOrderReq {
    1:optional i64 memberId=0 (api.raw = "memberId")
    2:optional i64 createdId=1 (api.raw = "createdId")
    3:optional Item items={} (api.raw = "items")
    4:optional double totalAmount=0 (api.raw = "totalAmount")
}

struct PaymentReq {

}

service OrderService  {
    OrderResp GetOrderInfo(1:GetOrderInfoReq req)
    GetOrderListResp GetOrderList(1:GetOrderListReq req)
    base.BaseResp DeleteOrder(1:base.IDReq req)
    OrderResp CreateOrder(1:GetOrderListReq req)
    OrderResp Payment(1:PaymentReq req)
    base.BaseResp CancelledOrder(1:CreateOrderReq req)
    base.BaseResp RefundOrder(1:RefundOrderReq req)
}