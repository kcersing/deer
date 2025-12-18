namespace go base

struct Api{
    1: optional i64 id=0 (api.raw = "id")
    2: optional string createdAt="" (api.raw = "createdAt")
    3: optional string updatedAt="" (api.raw = "updatedAt")
    4: optional string path="" (api.raw = "path")
    5: optional string desc="" (api.raw = "desc")
    6: optional string group="" (api.raw = "group")
    7: optional string method="" (api.raw = "method")
    8: optional string title="" (api.raw = "title")
   256:optional i64 createdId=0 (api.raw = "createdId" )
   257:optional string createdName="" (api.raw = "createdName" )
}

struct Menu{
    1:optional i64 id=0 (api.raw = "id" )
    2:optional string name="" (api.raw = "name" api.vd = "len($) > 0 && len($) < 33>")
    3:optional i64 parentId=0 (api.raw = "parentId")
    4:optional i64 level=0 (api.raw = "level")
    5:optional string path="" (api.raw = "path")
    6:optional string redirect="" (api.raw = "redirect")
    7:optional string component="" (api.raw = "component")
    8:optional i64 menuType=0 (api.raw = "menuType")
    12:optional i64 status=0 (api.raw = "status")
    14:optional list<Menu> children=[]  (api.raw = "children")
    15:optional string createdAt="" (api.raw = "createdAt")
    16:optional string updatedAt="" (api.raw = "updatedAt")
    20:optional string icon="" (api.raw = "icon")
   256:optional i64 createdId=0 (api.raw = "createdId" )
   257:optional string createdName="" (api.raw = "createdName" )
}

struct Role {
    1:optional i64 id=0 (api.raw = "id")
    2:optional string name="" (api.raw = "name")
    3:optional string code=""(api.raw = "code")
    4:optional string desc=""(api.raw = "desc")
    5:optional i64 orderNo=0(api.raw = "orderNo")
    6:optional list<i64> apis= [](api.raw = "apis")
    7:optional list<i64> menus= [](api.raw = "menus")
   256:optional i64 createdId=0 (api.raw = "createdId" )
   257:optional string createdName="" (api.raw = "createdName" )
}

// 字典信息
struct Dict {
    1:  i64 id=0 (api.raw = "id" )
    2:  string title="" (api.raw = "title" )
    3:  string code="" (api.raw = "code" )
    5:  i64 status=0 (api.raw = "status" )
    6:  string desc="" (api.raw = "desc" )
    7:  string createdAt="" (api.raw = "createdAt" )
    8:  string updatedAt="" (api.raw = "updatedAt" )
   256:optional i64 createdId=0 (api.raw = "createdId" )
   257:optional string createdName="" (api.raw = "createdName" )

}

// 字典键值信息
struct Dictht {
    1:  i64 id=0 (api.raw = "id" )
    2:  string title="" (api.raw = "title" )
    4:  string value="" (api.raw = "value" )
    5:  i64 status=0 (api.raw = "status" )
    6:  string createdAt="" (api.raw = "createdAt" )
    7:  string updatedAt="" (api.raw = "updatedAt" )
    8:  i64 dictId (api.raw = "dictId" )
   256:optional i64 createdId=0 (api.raw = "createdId" )
   257:optional string createdName="" (api.raw = "createdName" )
}

//日志信息
struct Log {
    1: optional string type="" (api.raw = "type")
    2: optional string method =""(api.raw = "method")
    3: optional string api =""(api.raw = "api")
    4: optional bool success = true (api.raw = "success")
    5: optional string reqContent="" (api.raw = "reqContent")
    6: optional string respContent="" (api.raw = "respContent")
    7: optional string ip="" (api.raw = "ip")
    8: optional string userAgent =""(api.raw = "userAgent")
    9: optional string operatorsr =""(api.raw = "operatorsr")
    10: optional i64 time=0 (api.raw = "time")
    11: optional string createdAt =""(api.raw = "createdAt")
    12: optional string updatedAt =""(api.raw = "updatedAt")
    13: optional i64 identity =0(api.raw = "identity")
    251: optional i64 id = 0 (api.raw = "id")
}