namespace go member
include "../base.thrift"

struct Member {
    254:optional i64 id=0 (api.raw = "id")
}
struct MemberResp {
    1:optional Member Member
    255:optional base.BaseResp baseResp
}

service MemberService  {
     MemberResp GetMemberInfo(1:base.IdReq req)
}