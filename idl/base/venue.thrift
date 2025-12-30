namespace go base
include "base.thrift"
struct Venue {
    1:  optional i64 id=0 (api.raw = "id")
    2:  optional string name="" (api.raw = "name")
    3:  optional string address="" (api.raw = "address")
    5:  optional string latitude="" (api.raw = "latitude")
    6:  optional string longitude =""(api.raw = "longitude")
    7:  optional string mobile="" (api.raw = "mobile")
    8:  optional string pic =""(api.raw = "pic")
    /**详情*/
    9:  optional string desc="" (api.raw = "desc")
     /**是否开放:0关闭 1开放;*/
    10:  optional i64 status=1 (api.raw = "status")
    11:  optional string createdAt="" (api.raw = "createdAt")
    12:  optional string updatedAt="" (api.raw = "updatedAt")
    13:  optional string email =""(api.raw = "email")

    14:  optional string startTime="" (api.raw = "startTime")
    15:  optional string endTime="" (api.raw = "endTime")
}
struct Place{
    1: optional i64 id=0 (api.raw = "id")
    2: optional string name="" (api.raw = "name")
    3: optional i64 venueId=0 (api.raw = "venueId")
    4: optional string pic="" (api.raw = "pic")
    /**是否开放:0关闭 1开放;*/
    5: optional i64 status=1 (api.raw = "status")
    6: optional string createdAt="" (api.raw = "createdAt")
    7: optional string updatedAt="" (api.raw = "updatedAt")
    /**可容纳人数*/
    8: optional i64 number=0 (api.raw = "number")
    /**详情*/
    9: optional string desc="" (api.raw = "desc")

    /**关联座位*/
    10: optional list<list<base.Seat>> seat=0 (api.raw = "seat")

    11:  optional string startTime="" (api.raw = "startTime")
    12:  optional string endTime="" (api.raw = "endTime")

}