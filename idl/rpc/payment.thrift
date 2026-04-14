namespace go payment
include "../base/base.thrift"


struct WXPayReq {
    1:optional i64 orderId=0 (api.raw = "orderId")
    2:optional string  orderSn="" (api.raw = "orderSn")
    3:optional double  total=0 (api.raw = "total")
    4:optional string  openId="" (api.raw = "openId")
    6:optional string  productName="" (api.raw = "productName")
}

struct WXQRPayReq {
    1: string id (api.raw = "id")
    2: string createTime (api.raw = "create_time")
    3: string resourceType (api.raw = "resource_type")
    4: string eventType (api.raw = "event_type")
    5: string summary (api.raw = "summary")
    6: Resource resource (api.raw = "resource")
}
struct WXRefundReq {
   1: string transactionId (api.raw = "transactionId")
    2: string outRefundNo (api.raw = "outRefundNo")
    3: double fee (api.raw = "fee")
    4: double total (api.raw = "total")
}

struct Resource {
    1: string original_type (api.raw = "original_type")
    2: string algorithm (api.raw = "algorithm")
    3: string ciphertext (api.raw = "ciphertext")
    4: string associated_data (api.raw = "associated_data")
    5: string nonce (api.raw = "nonce")
}
struct WXNotifyReq {

}
struct WXRefundNotifyReq {

}
service PaymentService {
    base.NilResponse WXPay(1: WXPayReq req)
    base.NilResponse WXQRPay(1: WXQRPayReq req)
    base.NilResponse WXRefund(1: WXRefundReq req)
}