namespace go system
include "../base/base.thrift"
include "../base/system.thrift"

struct ApiResp{
   1:optional system.Api data={}
   255:optional base.BaseResp baseResp={}
}
struct ApiListResp{
   1:optional list<system.Api> data={}
   255:optional base.BaseResp baseResp={}
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
struct CreateMenuReq{
    2:optional string name="" (api.raw = "name" api.vd = "len($) > 0 && len($) < 33>")
    3:optional i64 parentId=0 (api.raw = "parentId")
    4:optional i64 level=0 (api.raw = "level")
    5:optional string path="" (api.raw = "path")
    6:optional string redirect="" (api.raw = "redirect")
    7:optional string component="" (api.raw = "component")
    8:optional i64 menuType=0 (api.raw = "menuType")
    9:optional i64 orderNo=0 (api.raw = "orderNo")
    10:optional i64 ignore=0 (api.raw = "ignore")
    11:optional string icon="" (api.raw = "icon")
    12:optional i64 status=0 (api.raw = "status")
    15:optional string createdAt="" (api.raw = "createdAt")
    16:optional string updatedAt="" (api.raw = "updatedAt")

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
    9:optional i64 orderNo=0 (api.raw = "orderNo")
    10:optional i64 ignore=0 (api.raw = "ignore")
    11:optional string icon="" (api.raw = "icon")
    12:optional i64 status=0 (api.raw = "status")
    15:optional string createdAt="" (api.raw = "createdAt")
    16:optional string updatedAt="" (api.raw = "updatedAt")
}
struct MenuListReq{
    1:  optional i64 page=1 (api.raw = "page")
    2:  optional i64 pageSize=100 (api.raw = "pageSize")
}
struct MenuTree {
    1:optional system.Menu menuInfo={}
    4:optional list<MenuTree> children= []
    5:optional bool ignore=false
}

struct MenuResp{
   1:optional system.Menu data={}
   255:optional base.BaseResp baseResp={}
}
struct MenuListResp{
   1:optional list<system.Menu> data= []
   255:optional base.BaseResp baseResp={}
}

struct RoleResp {
    1:optional system.Role data={}
    255:optional base.BaseResp baseResp={}
}
struct RoleListResp {
    1:optional list<system.Role> data= []
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
    1:optional system.Dict data={}
    255:optional base.BaseResp baseResp={}
}
struct DicthtResp{
    1:optional system.Dictht data={}
    255:optional base.BaseResp baseResp={}
}
struct DictListResp{
    1:optional list<system.Dict> data={}
    255:optional base.BaseResp baseResp={}
}
struct DicthtListResp{
    1:optional list<system.Dictht> data={}
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


struct DeleteLogReq {
    1:  optional string startAt = "" (api.raw = "startAt")
    2:  optional string endAt = "" (api.raw = "endAt")
}
struct LogListResp{
    1:optional list<system.Log> data={}
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
    DictResp CreateDict(1: system.Dict req)
    // 更新字典信息
    DictResp UpdateDict(1: system.Dict req)
    // 删除字典信息
    base.NilResponse DeleteDict(1: base.IdReq req)
    // 获取字典列表
    DictListResp DictList(1: DictListReq req)
    // 创建字典键值信息
    DicthtResp CreateDictht(1: system.Dictht req)
    // 更新字典键值信息
    DicthtResp UpdateDictht(1: system.Dictht req)
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
