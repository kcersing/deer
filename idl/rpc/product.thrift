namespace go product
include "../base/base.thrift"
include "../base/product.thrift"




struct CreateItemReq {

    1:optional i64 id =0
   /**名称 */
    2:optional string name="" (api.raw = "name")
    /**主图 */
    3:optional string pic="" (api.raw = "pic")
    /**详情 */
    4:optional string desc="" (api.raw = "desc")

    /**时长 */
    5:optional i64 duration =0
    /**单次时长 */
    6:optional i64 length =0
    /**次数 */
    7:optional i64 count =0
    /**类型 */
    8:optional string type=""

    11: optional list<i64> tagId=0   (api.raw = "tagId")

    14: optional i64 createdId = 0 (api.raw = "createdId")


}

struct UpdateItemReq {

   /**名称 */
    2:optional string name="" (api.raw = "name")
    /**主图 */
    3:optional string pic="" (api.raw = "pic")
    /**详情 */
    4:optional string desc="" (api.raw = "desc")

    /**时长 */
    5:optional i64 duration =0
    /**单次时长 */
    6:optional i64 length =0
    /**次数 */
    7:optional i64 count =0
    /**类型 */
    8:optional string type=""

    11: optional list<i64> tagId=0   (api.raw = "tagId")

    14: optional i64 createdId = 0 (api.raw = "createdId")

}

struct CreateProductReq {
    /**名称 */
    2:optional string name="" (api.raw = "name")
    /**主图 */
    3:optional string pic="" (api.raw = "pic")
    /**详情 */
    4:optional string desc="" (api.raw = "desc")
    /**状态*/
    8:optional i64 status=0 (api.raw = "status")


    20: optional i64 createdId = 0 (api.raw = "createdId")


    /**价格 */
    6:optional double price=0 (api.raw = "price")
    /**库存 */
    7:optional i64 stock=0 (api.raw = "stock")
    10: optional list<i64> isSales=0 (api.raw = "isSales")
    13: optional string signSalesAt = "" (api.raw = "signSalesAt")
    14: optional string endSalesAt = "" (api.raw = "endSalesAt")

}
struct EditProductReq {

    1:optional i64 id=0 (api.raw = "id")

    /**名称 */
    2:optional string name="" (api.raw = "name")
    /**主图 */
    3:optional string pic="" (api.raw = "pic")
    /**详情 */
    4:optional string desc="" (api.raw = "desc")
    /**状态*/
    8:optional i64 status=0 (api.raw = "status")


    20: optional i64 createdId = 0 (api.raw = "createdId")


    /**价格 */
    6:optional double price=0 (api.raw = "price")
    /**库存 */
    7:optional i64 stock=0 (api.raw = "stock")
    10: optional list<i64> isSales=0 (api.raw = "isSales")
    13: optional string signSalesAt = "" (api.raw = "signSalesAt")
    14: optional string endSalesAt = "" (api.raw = "endSalesAt")

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
struct ItemListReq{
    1: i64 page=0 (api.raw = "page")
    2: i64 pageSize=10 (api.raw = "pageSize")
    3: optional string name (api.raw = "name")
    4: optional list<i64> status =0(api.raw = "status")
    5: optional list<i64> venueId =0  (api.raw = "venue")
    6: optional list<string> createdAt=0  (api.raw = "createdAt")
    7: optional string type ="" (api.raw = "type")
}

struct ProductResp{
    1:optional product.Product data= {}
    255:optional base.BaseResp baseResp={}
}
struct ItemResp{
    1:optional product.ProductItem data= {}
    255:optional base.BaseResp baseResp={}
}
struct ProductListResp{
    1:optional list<product.Product> data= []
    255:optional base.BaseResp baseResp={}
}
struct ItemListResp{
    1:optional list<product.ProductItem> data= []
    255:optional base.BaseResp baseResp={}
}
struct SearchProductReq{

}
struct DecrStockReq{

}

service ProductService  {

    ProductResp CreateProduct(1: CreateProductReq req) // 添加商品
    ProductResp UpdateProduct(1: EditProductReq req) // 编辑商品
    base.NilResponse  DeleteProduct(1: base.IdReq req) // 删除商品
    base.NilResponse  OnlineProduct(1: base.IdReq req) // 上架商品
    base.NilResponse  OfflineProduct(1: base.IdReq req) // 下架商品
    ProductResp GetProduct(1: base.IdReq req)
    ProductListResp SearchProduct(1: SearchProductReq req) // 搜索商品
    ProductListResp ProductList(1: ListReq req) // 商品列表
    base.NilResponse DecrStock(1: DecrStockReq req) // 扣减库存
    base.NilResponse DecrStockRevert(1: DecrStockReq req) // 库存返还


    // 添加属性
    ItemResp CreateItem(1: CreateItemReq req)
    // 编辑属性
    ItemResp UpdateItem(1: UpdateItemReq req)
    // 删除属性
    base.NilResponse DeleteItem(1: base.IdReq req)
    // 商品列表
    ItemListResp ItemList(1: ItemListReq req)


}