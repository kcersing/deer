namespace go user
include "../base.thrift"

struct User {
    1: i64 id,
    2: string username,
    3: string password,
    4: string avatar,
    5: string mobile,
    6: string name,
    7: i64 status,
    9: i64 gender,
    10: string birthday,

    11: string  lastAt,//最后一次登录时间
    12: string  lastIp,//最后一次登录ip

    251: string createdAt,
    252: string updatedAt,
    253: string createdId,
}
struct Role{
    1: i64 roleId
    2: i64 roleName
}

struct CheckUserReq {
    1: string username (vt.min_size = "1")
    2: string password (vt.min_size = "1")
    3: string captcha (vt.min_size = "1")
}

struct CreateUserReq{

}
struct GetUserListReq{

}

service UserService  {
     base.NilResponse CreateUser(1: CreateUserReq req)(api.post = "/service/user/create")
     base.NilResponse GetUser(1: base.IdReq req)(api.post = "/service/user")
     base.NilResponse LoginUser(1: CheckUserReq req) (api.post = "/service/user/login")
     base.NilResponse GetUserList(1: GetUserListReq req)(api.post = "/service/user/list")
}