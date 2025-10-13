namespace go member
include "../base/base.thrift"
struct Member {
    1:optional i64 id,
    2:optional string username,
    3:optional string password,
    4:optional string avatar,
    5:optional string mobile,
    6:optional string name,
    7:optional i64 status,
    8:optional i64 level,
    9:optional i64 gender,
    10:optional string birthday,

    11:optional string  lastAt,//最后一次登录时间
    12:optional string  lastIp,//最后一次登录ip

    251:optional string createdAt,
    252:optional string updatedAt,
    253:optional string createdId,

}

struct CreateMemberReq{

}
struct GetMemberListReq{
  1:optional base.PageReq pages
}
struct MemberResp {
    1:optional Member member
    255:optional base.BaseResp baseResp
}

struct MemberListResp {
    1:optional list<Member> member
    255:optional base.BaseResp baseResp
}

service MemberService  {
     MemberResp CreateMember(1: CreateMemberReq req)(api.post = "/service/member/create")
     MemberResp GetMember(1: base.IdReq req)(api.post = "/service/member")
     MemberResp LoginMember(1: base.CheckAccountReq req) (api.post = "/service/member/login")
     MemberListResp GetMemberList(1: GetMemberListReq req)(api.post = "/service/member/list")
}