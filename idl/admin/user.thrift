namespace go user
include "../base/base.thrift"


struct CreateUserReq{
     1:optional string username(api.raw = "username"),
     2:optional string password(api.raw = "password"),
}

struct UpdateUserReq {
    1:optional i64 id=0(api.raw = "id"),
    2:optional string avatar=""(api.raw = "avatar"),
    3:optional string mobile=""(api.raw = "mobile"),
    4:optional string name=""(api.raw = "name"),
    5:optional i64 status=0(api.raw = "status"),
    6:optional i64 gender=0(api.raw = "gender"),
    7:optional string birthday=""(api.raw = "birthday"),
    8:optional string  desc=""(api.raw = "desc"),
    9:optional i64 roleId=0(api.raw = "roleId"),
    10:optional i64 departmentsId=0(api.raw = "departmentsId"),
    11:optional i64 positionsId=0(api.raw = "positionsId"),


}
struct GetUserListReq{
      1:optional i64 page=1(api.raw = "page")
      2:optional i64 pageSize=10(api.raw = "pageSize")
      3:optional string keyword=""(api.raw = "keyword")
      4:optional string name=""(api.raw = "name")
      5:optional string mobile=""(api.raw = "mobile"),
}
struct ChangePasswordReq{
    1:optional i64 id=0(api.raw = "id"),
    2:optional string password=""(api.raw = "password"),
}
struct SetUserRoleReq{
   1:optional i64 id=0(api.raw = "id"),
    2:optional i64 roleId=0(api.raw = "roleId"),
}

service UserService  {
     base.NilResponse CreateUser(1: CreateUserReq req)(api.post = "/service/user/create")
     base.NilResponse GetUser(1: base.IdReq req)(api.post = "/service/user")
     base.NilResponse GetUserList(1: GetUserListReq req)(api.post = "/service/user/list")
     base.NilResponse UpdateUser(1: UpdateUserReq req)(api.post = "/service/user/update")
     base.NilResponse ChangePassword(1: ChangePasswordReq req)(api.post = "/service/user/change-password")
     base.NilResponse DeleteUser(1: base.IdReq req)(api.post = "/service/user/delete")
     base.NilResponse SetUserRole(1: SetUserRoleReq  req)(api.post = "/service/user/set/role")
}