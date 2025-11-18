namespace go positions
include "../base/base.thrift"

struct CreatePositionsReq{
     2:optional string name="",
     3:optional string code="",
     4:optional string departmentId="",
     5:optional string parentId="",
     6:optional string desc="",
     7:optional i64 status=0,
     9:optional i64 quota=0,
}
struct UpdatePositionsReq{
     1:optional i64 id=0,
     2:optional string name="",
     3:optional string code="",
     4:optional string departmentId="",
     5:optional string parentId="",
     6:optional string desc="",
     7:optional i64 status=0,
     9:optional i64 quota=0,
}
struct GetPositionsListReq{
      1:optional i64 page=1
      2:optional i64 pageSize=10
      3:optional string keyword=""
}
service UserService  {
    base.NilResponse CreatePositions(1: CreatePositionsReq req)(api.post = "/service/positions/create")
    base.NilResponse DeletePositions(1: base.IdReq req)(api.post = "/service/positions/delete")
    base.NilResponse UpdatePositions(1: UpdatePositionsReq req)(api.post = "/service/positions/update")
    base.NilResponse GetPositions(1: base.IdReq req)(api.post = "/service/positions")
    base.NilResponse GetPositionsList(1: GetPositionsListReq req)(api.post = "/service/positions/list")
}