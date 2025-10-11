namespace go member
include "../base.thrift"
struct Member {
    1: i64 id,
    2: string username,
    3: string password,
    4: string avatar,
    5: string mobile,
    6: string name,
    7: i64 status,
    8: i64 level,
    9: i64 gender,
    10: string birthday,

    11: string  lastAt,//最后一次登录时间
    12: string  lastIp,//最后一次登录ip

    251: string createdAt,
    252: string updatedAt,
    253: string createdId,
}
service MemberService  {
       base.NilResponse GetOrderById(1: base.IdReq req) (api.post = "/service/order/GetMemberInfo")
}