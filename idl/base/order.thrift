namespace go base

struct OrderItem {
    1:optional i64 productId=0 (api.raw = "productId")
    2:optional i64 quantity=1 (api.raw = "quantity")
    3:optional i64 price=0 (api.raw = "price")
    4:optional string name="" (api.raw = "name")
}
struct Order {
    1:optional i64 memberId=0 (api.raw = "memberId")
    2:optional i64 createdId=0 (api.raw = "createdId")

    3:optional list<OrderItem> items={} (api.raw = "items")
    4:optional string sn="" (api.raw = "sn")
    5:optional i64 totalAmount=0 (api.raw = "totalAmount")
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
    4:optional i64 refundAmount=0 (api.raw = "refundAmount")
}



struct OrderPay {
  1:optional i64  remission=0 (api.raw = "remission")
  2:optional i64  pay=0 (api.raw = "pay")
  3:optional string  reason="" (api.raw = "reason")
  4:optional string  payAt="" (api.raw = "payAt")
  5:optional string  payWay="" (api.raw = "payWay")
  6:optional string  paySn="" (api.raw = "paySn")
  7:optional string  prepayId="" (api.raw = "prepayId"),
  8:optional string  payExtra="" (api.raw = "payExtra")
  9:optional i64 createdId=0 (api.raw = "createdId")
}

