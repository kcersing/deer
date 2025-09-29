
namespace go base

struct BaseResp {
    1: i64 code=0
    2: string message=""
    3: string time=""
    4: i64 total=0
}
struct PageReq {
    1: i64 page=1
    2: i64 pageSize=10
    3: optional string Keyword=""
}
struct IdReq {
    1: i64 id
}
struct NilResponse {

}