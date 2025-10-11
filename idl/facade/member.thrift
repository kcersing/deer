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
struct CheckMemberReq {
    1: string username (vt.min_size = "1")
    2: string password (vt.min_size = "1")
    3: string captcha (vt.min_size = "1")
}

struct CreateMemberReq{

}
struct GetMemberListReq{

}
service MemberService  {
     base.NilResponse CreateMember(1: CreateMemberReq req)(api.post = "/service/member/create")
     base.NilResponse GetMember(1: base.IdReq req)(api.post = "/service/member")
     base.NilResponse LoginMember(1: CheckMemberReq req) (api.post = "/service/member/login")
     base.NilResponse GetMemberList(1: GetMemberListReq req)(api.post = "/service/member/list")
}