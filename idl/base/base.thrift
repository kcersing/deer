namespace go base

struct BaseResp {
    1:optional i64 code=0
    2:optional string message=""
    3:optional string time=""
    4:optional i64 total=0
}
struct PageReq {
    1:optional i64 page=1
    2:optional i64 pageSize=10
    3:optional string Keyword=""
}
struct IdReq {
    1:optional i64 id=0
}
struct NilResponse {

}
struct List{
     1:optional	i64 id =0 (api.raw = "id")
     2:optional	string title ="" (api.raw = "title")
}

struct CheckAccountReq {
    1:string username (api.raw = "username")
    2:string password (api.raw = "password")
    3:string captcha (api.raw = "captcha")
}
struct Tree  {
     1:optional	string title ="" (api.raw = "title")
     2:optional	string value="" (api.raw = "value")
     3:optional	string key="" (api.raw = "key")
     4:optional	string method="" (api.raw = "method")
     5:optional	list<Tree> children={} (api.raw = "children")
}
