namespace go member
include "../base.thrift"

service MemberService  {
       base.NilResponse GetOrderById(1: base.IdReq req) (api.post = "/service/order/GetMemberInfo")
}