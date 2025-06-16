namespace go order

struct Req {
    1:  optional i64 id=0 (api.raw = "id")
}

struct Resp {
    1:  optional i64 id=0 (api.raw = "id")
}


service order {
    Resp info(1: Req req)
}