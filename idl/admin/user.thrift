namespace go user
include "../base/base.thrift"
struct User {
    1:optional i64 id,
    2:optional string username,
    3:optional string password,
    4:optional string avatar,
    5:optional string mobile,
    6:optional string name,
    7:optional i64 status,
    9:optional i64 gender,
    10:optional string birthday,

    11:optional string  lastAt,//最后一次登录时间
    12:optional string  lastIp,//最后一次登录ip

    251:optional string createdAt,
    252:optional string updatedAt,
    253:optional string createdId,
}
struct Role{
    1:optional i64 roleId
    2:optional i64 roleName
}


struct CreateUserReq{

}
struct GetUserListReq{

}

service UserService  {
     base.NilResponse CreateUser(1: CreateUserReq req)(api.post = "/service/user/create")
     base.NilResponse GetUser(1: base.IdReq req)(api.post = "/service/user")
     base.NilResponse LoginUser(1: base.CheckAccountReq req) (api.post = "/service/user/login")
     base.NilResponse GetUserList(1: GetUserListReq req)(api.post = "/service/user/list")
}