namespace go base

struct Product {
    1:optional i64 id=0 (api.raw = "id")

    /**名称 */
    2:optional string name="" (api.raw = "name")
    /**主图 */
    3:optional string pic="" (api.raw = "pic")
    /**详情 */
    4:optional string desc="" (api.raw = "desc")
    /**状态*/
    8:optional i64 status=0 (api.raw = "status")
    18: optional string statusName="" (api.raw = "statusName")
    19:optional string code="" (api.raw = "code")
    20: optional i64 createdId = 0 (api.raw = "createdId")
    21: optional string createdName = "" (api.raw = "createdName")

    16: optional string createdAt = ""  (api.raw = "createdAt")
    17: optional string updatedAt = "" (api.raw = "updatedAt")
    /**价格 */
    6:optional double price=0 (api.raw = "price")
    /**库存 */
    7:optional i64 stock=0 (api.raw = "stock")
    10: optional list<i64> isSales=0 (api.raw = "isSales")
    13: optional string signSalesAt = "" (api.raw = "signSalesAt")
    14: optional string endSalesAt = "" (api.raw = "endSalesAt")


    15: optional list<Item> items = "" (api.raw = "items")
}

struct Item {
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
//    /**激活时间 */
//    9:optional string activeAt="" (api.raw = "activeAt")
//    /**到期时间 */
//    10:optional string expiredAt="" (api.raw = "expiredAt")
    11: optional list<i64> tagId=0   (api.raw = "tagId")
    /**价格 */
    12:optional i64 price=0 (api.raw = "price")
    13:optional string code="" (api.raw = "code")
    14:optional i64 status=0 (api.raw = "status")
    20: optional i64 createdId = 0 (api.raw = "createdId")
    21: optional string createdName = "" (api.raw = "createdName")

    16: optional string createdAt = ""  (api.raw = "createdAt")
    17: optional string updatedAt = "" (api.raw = "updatedAt")

}