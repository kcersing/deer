namespace go member
include "../base/base.thrift"
include "../base/member.thrift"

struct CreateMemberReq{
    1:optional string name="",
    2:optional string mobile="",
    3:optional i64 createdId=0,
    4:optional string birthday="",
    5:optional i64 intention=0,
    6:optional i64 gender=0,
    7:optional i64 status=0,
    8:optional string avatar="",
}
struct GetMemberListReq{
     1:optional i64 page=1
     2:optional i64 pageSize=10
     3:optional string keyword=""
     4:optional list<i64> tags=[]
}
struct MemberResp {
    1:optional member.Member data={}
    255:optional base.BaseResp baseResp={}
}

struct MemberListResp {
    1:optional list<member.Member> data={}
    255:optional base.BaseResp baseResp={}
}
struct MemberIdsResp {
    1:optional list<i64> data={}
    255:optional base.BaseResp baseResp={}
}

struct UpdateMemberReq {

      1:optional i64 id=0,
      4:optional string avatar="",
      5:optional string mobile="",
      6:optional string name="",
      7:optional i64 status=0,
      8:optional i64 intention=0,
      9:optional i64 gender=0,
      10:optional string birthday="",
      253:optional i64 createdId=0

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
     MemberIdsResp GetMemberIds(1: GetMemberListReq req)
}