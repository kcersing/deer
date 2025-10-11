namespace go order
include "../base.thrift"

service OrderService {

  base.NilResponse GetOrderInfo(1: base.IdReq req) (api.post = "/service/order/GetOrderInfo")

 }



