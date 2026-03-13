namespace go base
include "base.thrift"
struct Member {
    1:optional i64 id=0,
    2:optional string username="",
    3:optional string password="",
    4:optional string avatar="",
    5:optional string mobile="",
    6:optional string name="",
    7:optional i64 status=0,
    8:optional i64 level=0,
    9:optional i64 gender=0,
    10:optional string birthday="",
    13:optional i64 intention=0,

    11:optional string  lastAt=""//最后一次登录时间
    12:optional string  lastIp=""//最后一次登录ip


    251:optional string createdAt=""
    252:optional string updatedAt=""
   256:optional i64 createdId=0
   257:optional string createdName="" 

}



struct MemberProduct {
    1:optional i64 id=0 (api.raw = "id")
    2:optional string name="" (api.raw = "name")
    3: optional string sn = ""  (api.raw = "sn")

    8:optional i64 status=0 (api.raw = "status")
    18: optional string statusName="" (api.raw = "statusName")
    19:optional string code="" (api.raw = "code")
    20: optional i64 createdId = 0 (api.raw = "createdId")
    21: optional string createdName = "" (api.raw = "createdName")

    16: optional string createdAt = ""  (api.raw = "createdAt")
    17: optional string updatedAt = "" (api.raw = "updatedAt")

    6:optional double price=0 (api.raw = "price")

    15: optional list<MemberProductItem> items = [] (api.raw = "items")

    22: optional i64 orderId = 0  (api.raw = "orderId")
    23: optional i64 venueId = 0  (api.raw = "venueId")
    24: optional i64 productId = 0 (api.raw = "productId")
    25: optional i64 memberId = 0 (api.raw = "memberId")

}
struct MemberProductItem {
    1:optional i64 id =0
    2:optional string name=""
    /**时长 */
    5:optional i64 duration =0
    /**单次时长 */
    6:optional i64 length =0
    /**次数 */
    7:optional i64 count =0
    18: optional i64 countUsed = 0
    /**类型 */
    8:optional string type=""
    /**激活时间 */
    9:optional string activeAt=""
    /**到期时间 */
    10:optional string expiredAt=""
    11: optional list<base.List> tags=[]
    /**价格 */
    12:optional i64 price=0
    13:optional string code=""
    14:optional i64 status=0
    20: optional i64 createdId = 0
    21: optional string createdName = ""

    16: optional string createdAt = ""
    17: optional string updatedAt = ""

}