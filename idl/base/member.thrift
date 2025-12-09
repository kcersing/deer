namespace go base
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


    11:optional string  lastAt=""//最后一次登录时间
    12:optional string  lastIp=""//最后一次登录ip
    13:optional i64 intention=0,

    251:optional string createdAt=""
    252:optional string updatedAt=""
    253:optional i64 createdId=0

}
