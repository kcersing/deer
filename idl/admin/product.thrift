namespace go product
include "../base.thrift"

service ProductService  {
     base.NilResponse GetProductInfo(1: base.IdReq req) (api.post = "/service/product/GetProductInfo")
}