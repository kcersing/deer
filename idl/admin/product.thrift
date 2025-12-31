namespace go product
include "../base/base.thrift"


struct CreateProductReq{
    1:optional string code="" (api.raw = "code")
    /**名称 */
    2:optional string name="" (api.raw = "name")
    /**主图 */
    3:optional string pic="" (api.raw = "pic")
    /**详情 */
    4:optional string desc="" (api.raw = "desc")
    /**状态*/
    8:optional i64 status=0 (api.raw = "status")
    /**价格 */
    6:optional i64 price=0 (api.raw = "price")
    /**库存 */
    7:optional i64 stock=0 (api.raw = "stock")

    9: optional list<i64> items={} (api.raw = "items")
    10: optional list<i64> isSales={} (api.raw = "isSales")
    13: optional list<string> salesAt = [] (api.raw = "salesAt")

}
struct UpdateProductReq{
   1:optional i64 id=0 (api.raw = "id")

    /**名称 */
    2:optional string name="" (api.raw = "name")
    /**主图 */
    3:optional string pic="" (api.raw = "pic")
    /**详情 */
    4:optional string desc="" (api.raw = "desc")
    /**状态*/
    8:optional i64 status=0 (api.raw = "status")

    5:optional string code="" (api.raw = "code")


    /**价格 */
    6:optional i64 price=0 (api.raw = "price")
    /**库存 */
    7:optional i64 stock=0 (api.raw = "stock")
    9: optional list<i64> items={} (api.raw = "items")
    10: optional list<i64> isSales={} (api.raw = "isSales")
    13: optional list<string> salesAt = [] (api.raw = "salesAt")
}
struct GetProductListReq{
    1:optional i64 page=1
    2:optional i64 pageSize=10
    3:optional string keyword=""
}


struct CreateItemReq{
    /**名称 */
    2:optional string name=""
    /**主图 */
    3:optional string pic=""
    /**详情 */
    4:optional string desc=""

    /**时长 */
    5:optional i64 duration =0
    /**单次时长 */
    6:optional i64 length =0
    /**次数 */
    7:optional i64 count =0
    /**类型 */
    8:optional string type=""
    9:optional string code=""
    11: optional list<i64> tagId=0
    12:optional i64 status=0
    13:optional i64 price=0

}
struct UpdateItemReq{
       1:optional i64 id=0
       /**名称 */
        2:optional string name=""
        /**主图 */
        3:optional string pic=""
        /**详情 */
        4:optional string desc=""

        /**时长 */
        5:optional i64 duration =0
        /**单次时长 */
        6:optional i64 length =0
        /**次数 */
        7:optional i64 count =0
        /**类型 */
        8:optional string type=""
        9:optional string code=""
        11: optional list<i64> tagId=0
       12:optional i64 status=0
        13:optional i64 price=0
}
struct GetItemListReq{
    1:optional i64 page=1
    2:optional i64 pageSize=10
    3:optional string keyword=""
}
struct SearchProductReq{
    1:optional i64 page=1
    2:optional i64 pageSize=10
    3:optional string keyword=""
}
struct DecrStockReq{
    1:optional i64 productId=0,
     2:optional i64 count=0,
}


service ProductService  {

    base.NilResponse CreateProduct(1: CreateProductReq req)(api.post = "/service/product/create") // 添加商品
    base.NilResponse UpdateProduct(1: UpdateProductReq req)(api.post = "/service/product/update") // 编辑商品
    base.NilResponse DeleteProduct(1: base.IdReq req)(api.post = "/service/product/delete") // 删除商品
    base.NilResponse OnlineProduct(1: base.IdReq req)(api.post = "/service/product/online") // 上架商品
    base.NilResponse OfflineProduct(1: base.IdReq req)(api.post = "/service/product/offline") // 下架商品

    base.NilResponse GetProduct(1: base.IdReq req)(api.post = "/service/product")
    base.NilResponse SearchProduct(1: SearchProductReq req)(api.post = "/service/product/search") // 搜索商品
    base.NilResponse GetProductList(1: GetProductListReq req)(api.post = "/service/product/list") // 商品列表

    // 添加属性
    base.NilResponse CreateItem(1: CreateItemReq req)(api.post = "/service/item/create")
    // 编辑属性
    base.NilResponse UpdateItem(1: UpdateItemReq req)(api.post = "/service/item/update")
    // 删除属性
    base.NilResponse DeleteItem(1: base.IdReq req)(api.post = "/service/item/delete")
    // 商品列表
    base.NilResponse GetItemList(1: GetItemListReq req)(api.post = "/service/item/list")

    base.NilResponse GetItem(1: base.IdReq req)(api.post = "/service/item")


}