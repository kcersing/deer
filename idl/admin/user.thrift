namespace go user
include "../base/base.thrift"
struct User {
    1:optional i64 id=0,
    2:optional string username="",
    3:optional string password="",
    4:optional string avatar="",
    5:optional string mobile="",
    6:optional string name="",
    7:optional i64 status=1,
    8:optional i64 level=0,
    9:optional i64 gender=0,
    10:optional string birthday="",

    11:optional string  lastAt="",//最后一次登录时间
    12:optional string  lastIp="",//最后一次登录ip
    13:optional string  detail="",

    14:optional list<Role> roles={},//角色
    251:optional string createdAt="",
    252:optional string updatedAt="",
    253:optional i64 createdId=0,
}
struct Role{
    1:optional i64 id=0,
    2:optional string name="",
    3:optional string value="",
    4:optional string defaultRouter="",
    5:optional string remark="",
    6:optional list<i64> apis={},
}

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
    8:optional string  detail=""(api.raw = "detail"),
    9:optional i64 roleId=0(api.raw = "roleId"),
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