namespace go order
include "../base/base.thrift"
//

struct Item {
    1:optional i64 productId=0 (api.raw = "productId")
    2:optional i64 quantity=1 (api.raw = "quantity")
    3:optional double price=0 (api.raw = "price")
    4:optional string name="" (api.raw = "name")
}
struct Order {
    1:optional i64 memberId=0 (api.raw = "memberId")
    2:optional i64 createdId=0 (api.raw = "createdId")

    3:optional list<Item> items={} (api.raw = "items")
    4:optional string sn="" (api.raw = "sn")
    5:optional double totalAmount=0 (api.raw = "totalAmount")
    6:optional string status="" (api.raw = "status")
    7:optional string nature="" (api.raw = "nature")
    8:optional string createdAt="" (api.raw = "createdAt")
    9:optional string completionAt="" (api.raw = "completionAt")
    10:optional string closeAt="" (api.raw = "closeAt")
    11:optional string updatedAt="" (api.raw = "updatedAt")
    12:optional string cancelledReason="" (api.raw = "cancelledReason")
    13:optional list<OrderPay> orderPays={} (api.raw = "orderPays")
    14:optional OrderRefund orderRefund={} (api.raw = "orderRefund")

    254:optional i64 id=0 (api.raw = "id")
}

struct OrderRefund {
    1:optional string refundAt="" (api.raw = "refundAt")
    2:optional string refundReason="" (api.raw = "refundReason")
    3:optional i64 createdId=0 (api.raw = "createdId")
    4:optional double refundAmount=0 (api.raw = "refundAmount")
}



struct OrderPay {
  1:optional double  remission=0 (api.raw = "remission")
  2:optional double  pay=0 (api.raw = "pay")
  3:optional string  reason="" (api.raw = "reason")
  4:optional string  payAt="" (api.raw = "payAt")
  5:optional string  payWay="" (api.raw = "payWay")
  6:optional string  paySn="" (api.raw = "paySn")
  7:optional string  prepayId="" (api.raw = "prepayId"),
  8:optional string  payExtra="" (api.raw = "payExtra")
  9:optional i64 createdId=0 (api.raw = "createdId")
}

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
    4:optional i64 createdId=0 (api.raw = "createdId")
}

struct CancelledOrderReq {
    1:optional i64 id=0 (api.raw = "id")
    2:optional string reason="" (api.raw = "reason")
    3:optional i64 createdId=0 (api.raw = "createdId")
}

struct CreateOrderResp {
    1:optional Order data
    255:optional base.BaseResp baseResp
}

struct OrderResp {
    1:optional Order data
    255:optional base.BaseResp baseResp
}
struct CreateOrderReq {
    1:optional i64 memberId=0 (api.raw = "memberId")
    2:optional i64 createdId=0 (api.raw = "createdId")
    3:optional Item items={} (api.raw = "items")
    4:optional double totalAmount=0 (api.raw = "totalAmount")
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