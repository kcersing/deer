namespace go base
include "base.thrift"
struct Schedule{
	1:optional i64 id= 0
	2:optional string type = ""
    4:optional i64 venueId =0
    5:optional i64 placeId = 0
    6:optional i64 num = 0
    7:optional i64 numSurplus = 0
    9:optional string startTime = ""
    10:optional string endTime = ""
    11:optional i64 price = 0
    12:optional string name = ""

    14:optional i64 userId = 0
    15:optional list<i64> memberIds = {}
    /**状态 是 0未发布 1发布 2取消*/
    18:optional i64 status = 0
    19:optional i64 productId = 0
    25:optional list<list<base.Seat>> seats= {}
    31:optional i64 signNum = 0

    20:optional string venueName = ""
    21:optional string placeName = ""

    32: string statusName = ""
    33:optional string userName = ""
    34:optional list<string> memberNames = ""

    253:optional string date = ""
}
struct MemberSchedule{
	1:optional i64 id=0

    3:optional i64 venueId = 0
    4:optional i64 placeId = 0
    9:optional i64 userId = 0
    11:optional i64 productId = 0
    13:optional i64 scheduleId=0
    16:optional i64 memberId=0
    17:optional i64 memberProductId=0

	2:optional string type = ""
    7:optional i64 price = 0
    8:optional string name = ""

    10:optional i64 status = 0

    5:optional string startTime = ""
    6:optional string endTime = ""
	14:optional string signStartTime =""
	15:optional string signEndTime   =""

    30:optional string venueName = ""
    31:optional string placeName = ""
    32:optional string userName = ""
    33:optional string statusName = ""
    34:optional string memberName = ""
    12:optional string memberProductName =""

    253:optional string date=""

}
struct UserSchedule{
	1:optional i64 id=""
    9:optional i64 userId = 0
    11:optional i64 productId = 0
	13:optional i64 scheduleId=0
    3:optional i64 venueId = 0
    4:optional i64 placeId = 0
    16:optional i64 memberId=0
    12:optional i64 memberProductId = 0

    2:optional string type = ""
    7:optional i64 price = 0
    8:optional string name = ""
    10:optional i64 status = 0

    5:optional string startTime = ""
    6:optional string endTime = ""
	14:optional string signStartTime=""
	15:optional string signEndTime =""

    30:optional string venueName = ""
    31:optional string placeName = ""
    32:optional string userName = ""
    33:optional string statusName = ""
    34:optional string memberName = ""
    35:optional string memberProductName = ""

    253:optional string date=""
}