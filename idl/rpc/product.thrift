namespace go product
include "../base/base.thrift"

struct Product {
    1:optional i64 id=0 (api.raw = "id")
    2:optional Common common ={}
    3:optional Sales sales ={}

}
struct Property {
    1:optional i64 id =0
    2:optional Common common ={}
    3:optional Sales sales ={}
    4:optional Item item={}


}

struct Common{
    /**名称 */
    2:optional string name="" (api.raw = "name")
    /**主图 */
    3:optional string pic="" (api.raw = "pic")
    /**详情 */
    4:optional string description="" (api.raw = "description")
    /**状态*/
    8:optional i64 status=0 (api.raw = "status")
    18: optional string statusName="" (api.raw = "statusName")

    20: optional i64 createdId = 0 (api.raw = "createdId")
    21: optional string createdName = "" (api.raw = "createdName")

    16: optional string createdAt = ""  (api.raw = "createdAt")
    17: optional string updatedAt = "" (api.raw = "updatedAt")
}


struct Sales{
    /**价格 */
    6:optional double price=0 (api.raw = "price")
    /**库存 */
    7:optional i64 stock=0 (api.raw = "stock")
    10: optional list<i64> isSales=0 (api.raw = "isSales")
    13: optional string signSalesAt = "" (api.raw = "signSalesAt")
    14: optional string endSalesAt = "" (api.raw = "endSalesAt")


}

struct Item{
    /**时长 */
    4:optional i64 duration =0
    /**单次时长 */
    5:optional i64 length =0
    /**次数 */
    6:optional i64 count =0
    /**类型 */
    7:optional string type=""
    /**激活时间 */
    8:optional string activeAt="" (api.raw = "activeAt")
    /**到期时间 */
    9:optional string expiredAt="" (api.raw = "expiredAt")
    10: optional list<i64> tagId=0   (api.raw = "tagId")
}
struct CreatePropertyReq {
    /**名称*/
    2: optional string name ="" (api.raw = "name")
    /**定价*/
    3: optional double price =0(api.raw = "price")
    /**时长*/
    4: optional i64 duration =0(api.raw = "duration")
    /**单次时长*/
    5: optional i64 length =0(api.raw = "length")
    /**次数*/
    6: optional i64 count=0 (api.raw = "count")
    /**类型*/
    7: optional string type ="" (api.raw = "type")
    8: optional list<i64> venueId =0(api.raw = "venueId")
    15: optional i64 status =0 (api.raw = "status")
    /**标签-数组*/
    16: optional list<i64> tagId=0   (api.raw = "tagId")

    /**合同-数组*/
    17: optional list<i64> contractId =0 (api.raw = "contractId")


}

struct UpdatePropertyReq {
    1: optional i64 id =0 (api.raw = "id")
    /**名称*/
    2: optional string name ="" (api.raw = "name")
    /**定价*/
    3: optional double price =0(api.raw = "price")
    /**时长*/
    4: optional i64 duration =0(api.raw = "duration")
    /**单次时长*/
    5: optional i64 length =0(api.raw = "length")
    /**次数*/
    6: optional i64 count=0 (api.raw = "count")
    /**类型*/
    7: optional string type ="" (api.raw = "type")
    8: optional list<i64> venueId =0(api.raw = "venueId")
    15: optional i64 status =0 (api.raw = "status")
    /**标签-数组*/
    16: optional list<i64> tagId=0   (api.raw = "tagId")

    /**合同-数组*/
    17: optional list<i64> contractId =0 (api.raw = "contractId")


}

struct CreateProductReq {

    /**商品名*/
    2: optional string name="" (api.raw = "name")
    /**主图*/
    3: optional string pic="" (api.raw = "pic")
    /**详情*/
    4: optional string description ="" (api.raw = "description")
    /**价格*/
    5: optional double price =0(api.raw = "price")
    /**库存*/
    6: optional i64 stock =0(api.raw = "stock")
    /**属性*/
//    7:optional  list<i64> propertys =0(api.raw = "propertys")

    10: optional i64 CreateId=0 (api.raw = "createId")

    /**销售方式 1会员端*/
    12: optional i64 isSales = 0 (api.raw = "isSales")
    /**销售开始时间*/
    13: optional string signSalesAt ="" (api.raw = "signSalesAt")
    /**销售结束时间*/
    14: optional string endSalesAt ="" (api.raw = "endSalesAt")
    15: optional i64 status =0 (api.raw = "status")

}
struct EditProductReq {
    1: optional i64 id =0(api.raw = "id")
    /**商品名*/
    2: optional string name="" (api.raw = "name")
    /**主图*/
    3: optional string pic="" (api.raw = "pic")
    /**详情*/
    4: optional string description ="" (api.raw = "description")
    /**价格*/
    5: optional double price =0(api.raw = "price")
    /**库存*/
    6: optional i64 stock =0(api.raw = "stock")
    /**属性*/
    7:optional  list<i64> propertys =0(api.raw = "propertys")

    10: optional i64 CreateId=0 (api.raw = "createId")

    /**销售方式 1会员端*/
    12: optional i64 isSales = 0 (api.raw = "isSales")
    /**销售开始时间*/
    13: optional string signSalesAt ="" (api.raw = "signSalesAt")
    /**销售结束时间*/
    14: optional string endSalesAt ="" (api.raw = "endSalesAt")
    15: optional i64 status =0 (api.raw = "status")

}

struct ListReq {
    1: i64 page=0 (api.raw = "page")
    2: i64 pageSize =10(api.raw = "pageSize")
    3: optional string name="" (api.raw = "name")
    4: optional list<i64> status =0 (api.raw = "status")
    5: optional list<i64> venueId =0  (api.raw = "venue")
    6: optional list<string> createdAt=0  (api.raw = "createdAt")
    7: optional string type ="" (api.raw = "type")
}
struct PropertyListReq{
    1: i64 page=0 (api.raw = "page")
    2: i64 pageSize=10 (api.raw = "pageSize")
    3: optional string name (api.raw = "name")
    4: optional list<i64> status =0(api.raw = "status")
    5: optional list<i64> venueId =0  (api.raw = "venue")
    6: optional list<string> createdAt=0  (api.raw = "createdAt")
    7: optional string type ="" (api.raw = "type")
}

struct ProductResp{
    1:optional Product data= []
    255:optional base.BaseResp baseResp={}
}
struct PropertyResp{
    1:optional Property data= []
    255:optional base.BaseResp baseResp={}
}
struct ProductListResp{
    1:optional list<Product> data= []
    255:optional base.BaseResp baseResp={}
}
struct PropertyListResp{
    1:optional list<Property> data= []
    255:optional base.BaseResp baseResp={}
}

service ProductService  {



    ProductResp CreateProduct(1: CreateProductReq req) // 添加商品
    ProductResp UpdateProduct(1: EditProductReq req) // 编辑商品
    base.NilResponse  DeleteProduct(1: base.IdReq req) // 删除商品
    base.NilResponse  Online(1: base.IdReq req) // 上架商品
    base.NilResponse  Offline(1: base.IdReq req) // 下架商品
    ProductResp GetProduct(1: base.IdReq req)
    ProductListResp SearchProduct(1: SearchProductReq req) // 搜索商品
    ProductListResp ProductList(1: ListReq req) // 商品列表
     base.NilResponse DecrStock(1: DecrStockReq req) // 扣减库存
     base.NilResponse DecrStockRevert(1: DecrStockReq req) // 库存返还




    // 添加属性
    base.NilResponse CreateProperty(1: CreatePropertyReq req)
    // 编辑属性
    base.NilResponse UpdateProperty(1: UpdatePropertyReq req)
    // 删除属性
    base.NilResponse DeleteProperty(1: base.IdReq req)
    // 商品列表
    PropertyListResp PropertyList(1: PropertyListReq req)








    // 添加商品
    base.NilResponse Create(1: CreateOrUpdateReq req) (api.post = "/service/product/create")
    // 编辑商品
    base.NilResponse Update(1: CreateOrUpdateReq req) (api.post = "/service/product/update")
    // 删除商品
    base.NilResponse Delete(1: base.IdReq req) (api.post = "/service/product/del")
    // 商品列表
    base.NilResponse List(1: ListReq req) (api.post = "/service/product/list")
    // 上架0/下架1
    base.NilResponse UpdateStatus(1: base.StatusCodeReq req) (api.post = "/service/product/status")

    // 商品详情
    base.NilResponse InfoById(1: base.IdReq req) (api.post = "/service/product/info")

    base.NilResponse ProductListExport(1: ListReq req) (api.post = "/service/product/export")
}