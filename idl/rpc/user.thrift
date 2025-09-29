namespace go user
include "../base.thrift"

struct User {
    254:optional i64 id=0 (api.raw = "id")
}
struct UserResp {
    1:optional User user
    255:optional base.BaseResp baseResp
}

service UserService  {
     UserResp GetUserInfo(1:base.IdReq req)
}