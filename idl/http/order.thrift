namespace go order
include "../base.thrift"

// hz update -idl ../../idl/http/order.thrift -model_dir http_gen/model/  --unset_omitempty

struct GetOrderInfoReq {
    1:optional i64 id=0 (api.raw = "id")
    2:optional string sn (api.raw = "sn")
}
service OrderService {

  base.NilResponse GetOrderById(1: GetOrderInfoReq req) (api.post = "/service/order/GetOrderById")

 }



