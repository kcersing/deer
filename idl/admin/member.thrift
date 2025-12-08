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
struct UpdateMemberReq {
  1:optional i64 id=0,
    4:optional string avatar="",
    5:optional string mobile="",
    6:optional string name="",
    7:optional i64 status=0,
    8:optional i64 level=0,
    9:optional i64 gender=0,
    10:optional string birthday="",

}
struct ChangePasswordReq {
    1:optional i64 id=0,
    2:optional string password="",
}
service MemberService {


     base.NilResponse CreateMember(1: CreateMemberReq req) (api.post = "/service/member/create")
     base.NilResponse DeleteMember(1: base.IdReq req) (api.post = "/service/member/delete")
     base.NilResponse UpdateMember(1: UpdateMemberReq req) (api.post = "/service/member/update")
     base.NilResponse GetMember(1: base.IdReq req) (api.post = "/service/member")
     base.NilResponse GetMemberList(1: GetMemberListReq req) (api.post = "/service/member/list")
     base.NilResponse ChangePassword(1: ChangePasswordReq req) (api.post = "/service/member/change-password")

}