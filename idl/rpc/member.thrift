namespace go member
include "../base/base.thrift"
include "../base/member.thrift"

struct CreateMemberReq{
    1:optional string username="",
    2:optional string password="",
}
struct GetMemberListReq{
     1:optional i64 page=1
     2:optional i64 pageSize=10
     3:optional string keyword=""
}
struct MemberResp {
    1:optional member.Member data={}
    255:optional base.BaseResp baseResp={}
}

struct MemberListResp {
    1:optional list<member.Member> data={}
    255:optional base.BaseResp baseResp={}
}
struct UpdateMemberReq {
    1:optional member.Member data={}
    255:optional base.BaseResp baseResp={}
}
struct ChangePasswordReq {
    1:optional i64 id=0,
    2:optional string password="",
}
service MemberService  {

     MemberResp CreateMember(1: CreateMemberReq req)
     base.NilResponse DeleteMember(1: base.IdReq req)
     MemberResp UpdateMember(1: UpdateMemberReq req)
     MemberResp GetMember(1: base.IdReq req)
     MemberListResp GetMemberList(1: GetMemberListReq req)

     MemberResp LoginMember(1: base.CheckAccountReq req)
     base.NilResponse ChangePassword(1: ChangePasswordReq req)

}