namespace go product
include "../base/base.thrift"
include "../base/product.thrift"

struct CreateItemReq {

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
    10:optional i64 price=0
    11: optional list<i64> tagId=0
    14: optional i64 createdId = 0
    15:optional i64 status=0

}

struct UpdateItemReq {
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
    10:optional i64 price=0
    11: optional list<i64> tagId=0 
    14: optional i64 createdId = 0
    15:optional i64 status=0
}

struct CreateProductReq {
    1:optional string code=""
    /**名称 */
    2:optional string name="" 
    /**主图 */
    3:optional string pic="" 
    /**详情 */
    4:optional string desc=""
    /**状态*/
    8:optional i64 status=0 

    20: optional i64 createdId = 0 

    /**价格 */
    6:optional i64 price=0 
    /**库存 */
    7:optional i64 stock=0
    10: optional list<i64> isSales=0 
    13: optional string signSalesAt = "" 
    14: optional string endSalesAt = "" 

}
struct UpdateProductReq {

    1:optional i64 id=0

    /**名称 */
    2:optional string name=""
    /**主图 */
    3:optional string pic="" 
    /**详情 */
    4:optional string desc="" 
    /**状态*/
    8:optional i64 status=0 

    5:optional string code=""
    20: optional i64 createdId = 0

    /**价格 */
    6:optional i64 price=0 
    /**库存 */
    7:optional i64 stock=0
    10: optional list<i64> isSales=0 
    13: optional string signSalesAt = "" 
    14: optional string endSalesAt = "" 

}

struct ListReq {
    1: i64 page=0 
    2: i64 pageSize =10
    3: optional string name="" 
    4: optional list<i64> status =0 

}
struct ItemListReq{
    1: i64 page=0 
    2: i64 pageSize=10 
    3: optional string name
    4: optional list<i64> status =0
    7: optional string type =""
}

struct ProductResp{
    1:optional product.Product data= {}
    255:optional base.BaseResp baseResp={}
}
struct ItemResp{
    1:optional product.Item data= {}
    255:optional base.BaseResp baseResp={}
}
struct ProductListResp{
    1:optional list<product.Product> data= []
    255:optional base.BaseResp baseResp={}
}
struct ItemListResp{
    1:optional list<product.Item> data= []
    255:optional base.BaseResp baseResp={}
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

    ProductResp CreateProduct(1: CreateProductReq req) // 添加商品
    ProductResp UpdateProduct(1: UpdateProductReq req) // 编辑商品
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

    ItemResp GetItem(1: base.IdReq req)

}