namespace go user
include "../base.thrift"


service UserService  {
     base.NilResponse GetUserInfo(1: base.IdReq req) (api.post = "/service/user/GetUserInfo")
}