namespace go system
include "../base/base.thrift"
//日志列表请求数据
struct LogListReq {
    1: optional i64 page=0 (api.raw = "page")
    2: optional i64 pageSize=0 (api.raw = "pageSize")
    3: optional string type="" (api.raw = "type")
    4: optional string method="" (api.raw = "method")
    5: optional string api="" (api.raw = "api")
    6: optional bool success= true (api.raw = "success")
    7: optional string operatorsr ="" (api.raw = "operatorsr")
    8: optional i64 identity =0(api.raw = "identity")
}
struct DeleteLogReq {
    1:  optional string startAt = "" (api.raw = "startAt")
    2:  optional string endAt = "" (api.raw = "endAt")
}

struct CreateApiReq{
    4: optional string path="" (api.raw = "path")
    5: optional string desc="" (api.raw = "desc")
    6: optional string group="" (api.raw = "group")
    7: optional string method="" (api.raw = "method")
    8: optional string title="" (api.raw = "title")

}
struct UpdateApiReq{
    1: optional i64 id=0 (api.raw = "id")
    4: optional string path="" (api.raw = "path")
    5: optional string desc="" (api.raw = "desc")
    6: optional string group="" (api.raw = "group")
    7: optional string method="" (api.raw = "method")
    8: optional string title="" (api.raw = "title")
}
struct ApiListReq{
    3:optional string path = "" (api.raw = "path")
    4:optional string desc = "" (api.raw = "desc")
    5:optional string method  = ""(api.raw = "method")
    6:optional string group  = ""(api.raw = "group")


    251:optional i64 page=1 (api.raw = "page")
    252:optional i64 pageSize=10 (api.raw = "pageSize")
    253:optional string keyword="" (api.raw = "keyword")
}


service SystemService  {
    base.NilResponse CreateApi(1: CreateApiReq req)(api.post = "/service/api/create")
    base.NilResponse UpdateApi(1: UpdateApiReq req)(api.post = "/service/api/update")
    base.NilResponse DeleteApi(1: base.IdReq req)(api.post = "/service/api/delete")
    base.NilResponse ApiList(1: ApiListReq req) (api.post = "/service/api/list")
    base.NilResponse ApiTree(1: ApiListReq req)(api.post = "/service/api/tree")



    // Get logs list | 获取日志列表
    base.NilResponse LogList(1: LogListReq req) (api.post = "/service/logs/list")

    // Delete logs | 删除日志信息
    base.NilResponse DeleteLog(1: DeleteLogReq req) (api.post = "/service/logs/deleteAll")



}