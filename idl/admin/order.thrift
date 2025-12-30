namespace go order
include "../base/base.thrift"

service OrderService {

  base.NilResponse GetOrder(1: base.IdReq req) (api.post = "/service/order/GetOrderInfo")




 }



