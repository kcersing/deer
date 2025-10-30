namespace go system
include "../base/base.thrift"

struct Api{
    1: optional i64 id=0 (api.raw = "id")
    2: optional string createdAt="" (api.raw = "createdAt")
    3: optional string updatedAt="" (api.raw = "updatedAt")
    4: optional string path="" (api.raw = "path")
    5: optional string description="" (api.raw = "description")
    6: optional string group="" (api.raw = "group")
    7: optional string method="" (api.raw = "method")
    8: optional string title="" (api.raw = "title")
}
struct ApiResp{
   1:optional Api data={}
   255:optional base.BaseResp baseResp={}
}
struct ApiListResp{
   1:optional list<Api> data={}
   255:optional base.BaseResp baseResp={}
}
struct ApiListReq{
    3:optional string path = "" (api.raw = "path")
    4:optional string description = "" (api.raw = "description")
    5:optional string method  = ""(api.raw = "method")
    6:optional string group  = ""(api.raw = "group")


    251:optional i64 page=1 (api.raw = "page")
    252:optional i64 pageSize=10 (api.raw = "pageSize")
    253:optional string keyword="" (api.raw = "keyword")
}
struct CreateApiReq{
    4: optional string path="" (api.raw = "path")
    5: optional string description="" (api.raw = "description")
    6: optional string group="" (api.raw = "group")
    7: optional string method="" (api.raw = "method")
    8: optional string title="" (api.raw = "title")

}
struct UpdateApiReq{
    1: optional i64 id=0 (api.raw = "id")
    4: optional string path="" (api.raw = "path")
    5: optional string description="" (api.raw = "description")
    6: optional string group="" (api.raw = "group")
    7: optional string method="" (api.raw = "method")
    8: optional string title="" (api.raw = "title")
}
struct CreateMenuReq{
    2:optional string name="" (api.raw = "name" api.vd = "len($) > 0 && len($) < 33>")
    3:optional i64 parentId=0 (api.raw = "parentId")
    4:optional i64 level=0 (api.raw = "level")
    5:optional string path="" (api.raw = "path")
    6:optional string redirect="" (api.raw = "redirect")
    7:optional string component="" (api.raw = "component")
    8:optional i64 menuType=0 (api.raw = "menuType")
    9:optional i64 hidden=0 (api.raw = "hidden")
    10:optional i64 sort=0 (api.raw = "sort")

    12:optional i64 status=1 (api.raw = "status")
    13:optional string url="" (api.raw = "url")
//    14:optional list<Menu> children=[]  (api.raw = "children")
    15:optional string createdAt="" (api.raw = "createdAt")
    16:optional string updatedAt="" (api.raw = "updatedAt")
    17:optional string title="" (api.raw = "title" )
    19:optional string type="" (api.raw = "type")
}
struct UpdateMenuReq{
    1:optional i64 id=0 (api.raw = "id" )
    2:optional string name="" (api.raw = "name" api.vd = "len($) > 0 && len($) < 33>")
    3:optional i64 parentId=0 (api.raw = "parentId")
    4:optional i64 level=0 (api.raw = "level")
    5:optional string path="" (api.raw = "path")
    6:optional string redirect="" (api.raw = "redirect")
    7:optional string component="" (api.raw = "component")
    8:optional i64 menuType=0 (api.raw = "menuType")
    9:optional i64 hidden=0 (api.raw = "hidden")
    10:optional i64 sort=0 (api.raw = "sort")

    12:optional i64 status=1 (api.raw = "status")
    13:optional string url="" (api.raw = "url")
//    14:optional list<Menu> children=[]  (api.raw = "children")
    15:optional string createdAt="" (api.raw = "createdAt")
    16:optional string updatedAt="" (api.raw = "updatedAt")
    17:optional string title="" (api.raw = "title" )
    19:optional string type="" (api.raw = "type")
}
struct MenuListReq{
    1:  optional i64 page=1 (api.raw = "page")
    2:  optional i64 pageSize=100 (api.raw = "pageSize")
}
struct MenuTree {
    1:optional Menu menuInfo={}
    4:optional list<MenuTree> children= []
    5:optional bool ignore=false
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
    9:optional i64 hidden=0 (api.raw = "hidden")
    10:optional i64 sort=0 (api.raw = "sort")

    12:optional i64 status=1 (api.raw = "status")
    13:optional string url="" (api.raw = "url")
    14:optional list<Menu> children=[]  (api.raw = "children")
    15:optional string createdAt="" (api.raw = "createdAt")
    16:optional string updatedAt="" (api.raw = "updatedAt")
    17:optional string title="" (api.raw = "title" )
    19:optional string type="" (api.raw = "type")
    20:optional string icon="" (api.raw = "icon")
}
struct MenuResp{
   1:optional Menu data={}
   255:optional base.BaseResp baseResp={}
}
struct MenuListResp{
   1:optional list<Menu> data= []
   255:optional base.BaseResp baseResp={}
}
struct Role {
    1:optional i64 id=0 (api.raw = "id")
    2:optional string name="" (api.raw = "name")
    3:optional string value=""(api.raw = "value")
    4:optional string defaultRouter=""(api.raw = "defaultRouter")
    5:optional string remark=""(api.raw = "remark")
    6:optional list<i64> apis= [](api.raw = "apis")
}
struct RoleResp {
    1:optional Role data={}
    255:optional base.BaseResp baseResp={}
}
struct RoleListResp {
    1:optional list<Role> data= []
    255:optional base.BaseResp baseResp={}
}
struct CreateRoleReq{
    1:optional i64 id=0(api.raw = "id")
    2:optional string name=""(api.raw = "name")
    3:optional string value=""(api.raw = "value")
    4:optional string defaultRouter=""(api.raw = "defaultRouter")
    5:optional string remark=""(api.raw = "remark")
    6:optional list<i64> apis= [](api.raw = "apis")
}

struct GetRoleListReq{
    1:optional i64 page=1(api.raw = "page")
    2:optional i64 pageSize=10(api.raw = "pageSize")
    3:optional string keyword=""(api.raw = "keyword")
}
struct UpdateRoleReq {
    1:optional i64 id=0(api.raw = "id")
    2:optional string name=""(api.raw = "name")
    3:optional string value=""(api.raw = "value")
    4:optional string defaultRouter=""(api.raw = "defaultRouter")
    5:optional string remark=""(api.raw = "remark")
    6:optional list<i64> apis= [](api.raw = "apis")
}
struct CreateMenuAuthReq{
    1:optional i64 roleId=0 (api.raw = "roleId")
    2:optional list<i64> Ids = [] (api.raw = "ids")
}

// 字典信息
struct Dict {
    1:  i64 id=0 (api.raw = "id" )
    2:  string title="" (api.raw = "title" )
    3:  string name="" (api.raw = "name" )
    5:  i64 status=1 (api.raw = "status" )
    6:  string description="" (api.raw = "description" )
    7:  string createdAt="" (api.raw = "createdAt" )
    8:  string updatedAt="" (api.raw = "updatedAt" )
}

// 字典键值信息
struct Dictht {
    1:  i64 id=0 (api.raw = "id" )
    2:  string title="" (api.raw = "title" )
    3:  string key="" (api.raw = "key" )
    4:  string value="" (api.raw = "value" )
    5:  i64 status=1 (api.raw = "status" )
    6:  string createdAt="" (api.raw = "createdAt" )
    7:  string updatedAt="" (api.raw = "updatedAt" )
    8:  i64 parentID (api.raw = "parentID" )
}
// 字典列表请求数据
struct DictListReq {
    1:  optional string title  = ""(api.raw = "title" )
    2:  optional string name = "" (api.raw = "name" )
    3:  optional i64 page=1 (api.raw = "page")
    4:  optional i64 pageSize=100 (api.raw = "pageSize")
}

//字典名获取字典键值请求数据
struct DicthtListReq{
    1:  optional string name = "" (api.raw = "name" )
    2:  optional i64 dictionaryId= 0 (api.raw = "dictionaryId" )
}

struct DictResp{
    1:optional Dict data={}
    255:optional base.BaseResp baseResp={}
}
struct DicthtResp{
    1:optional Dictht data={}
    255:optional base.BaseResp baseResp={}
}
struct DictListResp{
    1:optional list<Dict> data={}
    255:optional base.BaseResp baseResp={}
}
struct DicthtListResp{
    1:optional list<Dictht> data={}
    255:optional base.BaseResp baseResp={}
}

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
struct DeleteLogReq {
    1:  optional string startAt = "" (api.raw = "startAt")
    2:  optional string endAt = "" (api.raw = "endAt")
}
struct LogListResp{
    1:optional list<Log> data={}
    255:optional base.BaseResp baseResp={}
}
struct TreeResp{
    1:optional list<base.Tree> data={}
    255:optional base.BaseResp baseResp={}
}
service SystemService  {
    ApiResp CreateApi(1: CreateApiReq req)
    ApiResp UpdateApi(1: UpdateApiReq req)
    base.NilResponse DeleteApi(1: base.IdReq req)
    ApiListResp ApiList(1: ApiListReq req)
    TreeResp ApiTree(1: ApiListReq req)

    MenuResp CreateMenu(1: CreateMenuReq req)
    MenuResp UpdateMenu(1: UpdateMenuReq req)
    base.NilResponse DeleteMenu(1: base.IdReq req)
    MenuResp GetMenu(1: base.IdReq req)
    MenuListResp MenuList(1: MenuListReq req)
    TreeResp MenuTree(1: MenuListReq req)

    RoleResp CreateRole(1: CreateRoleReq req)
    RoleResp GetRole(1: base.IdReq req)
    RoleListResp GetRoleList(1: GetRoleListReq req)
    RoleResp UpdateRole(1: UpdateRoleReq req)

    base.NilResponse DeleteRole(1: base.IdReq req)
    // 创建菜单权限
    base.NilResponse CreateRoleMenu(1: CreateMenuAuthReq req)
    // 创建API权限
    base.NilResponse CreateRoleApi(1: CreateMenuAuthReq req)
    // 获取角色API列表
    ApiListResp GetRoleApi(1: base.IdReq req)
    //获取角色菜单列表
    MenuListResp GetRoleMenu(1: base.IdReq req)

    // 创建字典信息
    DictResp CreateDict(1: Dict req)
    // 更新字典信息
    DictResp UpdateDict(1: Dict req)
    // 删除字典信息
    base.NilResponse DeleteDict(1: base.IdReq req)
    // 获取字典列表
    DictListResp DictList(1: DictListReq req)
    // 创建字典键值信息
    DictResp CreateDictht(1: Dictht req)
    // 更新字典键值信息
    DictResp UpdateDictht(1: Dictht req)
    // 删除字典键值信息
    base.NilResponse DeleteDictht(1: base.IdReq req)
    // 根据字典名获取字典键值列表
    DicthtListResp DicthtList(1: DicthtListReq req)

    // Get logs list | 获取日志列表
    LogListResp LogList(1: LogListReq req)

    // Delete logs | 删除日志信息
    base.NilResponse DeleteLog(1: DeleteLogReq req)

    // 验证权限
    VerifyRoleAuthResp VerifyRoleAuth(1: VerifyRoleAuthReq req)

}

struct VerifyRoleAuthReq{
  1: optional string obj="" (api.raw = "obj")
  2: optional string act =""(api.raw = "act")
  3: optional i64 roleId =0(api.raw = "roleId")
}
struct VerifyRoleAuthResp{
  255:optional base.BaseResp baseResp={}
 }
