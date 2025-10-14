namespace go product
include "../base/base.thrift"

struct Product {
    254:optional i64 id=0 (api.raw = "id")
}
struct ProductResp {
    1:optional Product data
    255:optional base.BaseResp baseResp
}

service ProductService  {
     ProductResp GetProductInfo(1:base.IdReq req)
}