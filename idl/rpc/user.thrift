namespace go user
include "../base/base.thrift"
include "../base/user.thrift"


struct UserResp {
    1:optional user.User data={}
    255:optional base.BaseResp baseResp={}
}
struct UserListResp {
    1:optional list<user.User> data= []
    255:optional base.BaseResp baseResp={}
}
struct CreateUserReq{
     1: string username,
     2: string password,
}

struct GetUserListReq{
      1:optional i64 page=1
      2:optional i64 pageSize=10
      3:optional string keyword=""
      4:optional string name=""
      5:optional string mobile="",
}
struct UpdateUserReq {
    1:optional i64 id=0,
    2:optional string avatar="",
    3:optional string mobile="",
    4:optional string name="",
    5:optional i64 status=0,
    6:optional i64 gender=0,
    7:optional string birthday="",
    8:optional string  detail="",
    9:optional i64 roleId=0,
}

struct ChangePasswordReq {
    1:optional i64 id=0,
    2:optional string password="",
}
struct SetUserRoleReq{
    1:optional i64 id=0,
    2:optional i64 roleId=0,
}

service UserService  {
    UserResp CreateUser(1: CreateUserReq req)
    UserResp GetUser(1: base.IdReq req)
    UserResp LoginUser(1: base.CheckAccountReq req)
    UserListResp GetUserList(1: GetUserListReq req)
    UserResp UpdateUser(1: UpdateUserReq req)
    
    base.NilResponse ChangePassword(1: ChangePasswordReq req)
    base.NilResponse DeleteUser(1: base.IdReq req)
    base.NilResponse SetUserRole(1: SetUserRoleReq  req)

}
