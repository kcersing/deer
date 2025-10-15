namespace go System
include "../base/base.thrift"

struct Api{
    1: optional i64 id (api.raw = "id")
    2: optional string createdAt (api.raw = "createdAt")
    3: optional string updatedAt (api.raw = "updatedAt")
    4: optional string path (api.raw = "path")
    5: optional string description (api.raw = "description")
    6: optional string group (api.raw = "group")
    7: optional string method (api.raw = "method")
}
struct ApiResp{
   1:optional Api data
   255:optional base.BaseResp baseResp
}
struct ApiListResp{
   1:optional Api data
   255:optional base.BaseResp baseResp
}
struct ApiListReq{
    1:optional base.PageReq pages
    3: string path = "" (api.raw = "path")
    4: string description = "" (api.raw = "description")
    5: string method  = ""(api.raw = "method")
    6: string group  = ""(api.raw = "group")
}
struct CreateApiReq{
    1:optional i64 id =0(api.raw = "id" )
    2:optional string name="" (api.raw = "name" api.vd = "len($) > 0 && len($) < 33>")
    3:optional i64 parentId=0 (api.raw = "parentId")
    4:optional i64 level=0 (api.raw = "level")
    5:optional string path="" (api.raw = "path")
    6:optional string redirect="" (api.raw = "redirect")
    7:optional string component="" (api.raw = "component")
    8:optional i64 menuType=0 (api.raw = "menuType")
    9:optional bool hidden=true (api.raw = "hidden")
    10:optional i64 sort=0 (api.raw = "sort")
    12:optional i64 status=1 (api.raw = "status")
    13:optional string url="" (api.raw = "url")
    14:optional string type="" (api.raw = "type")
}
struct UpdateApiReq{

}
struct CreateMenuReq{

}
struct UpdateMenuReq{

}
struct MenuListReq{

}
struct MenuTree {
    1: Menu menuInfo;
    4: list<MenuTree> children;
    5: bool ignore;
}
struct Menu{
    1: i64 id (api.raw = "id" )
    2: string name (api.raw = "name" api.vd = "len($) > 0 && len($) < 33>")
    3: i64 parentId (api.raw = "parentId")
    4: i64 level (api.raw = "level")
    5: string path (api.raw = "path")
    6: string redirect (api.raw = "redirect")
    7: string component (api.raw = "component")
    8: i64 menuType (api.raw = "menuType")
    9: bool hidden (api.raw = "hidden")
    10: i64 sort (api.raw = "sort")

    12: i64 status (api.raw = "status")
    13: string url (api.raw = "url")
    14: list<Menu> children  (api.raw = "children")
    15: string createdAt (api.raw = "createdAt")
    16: string updatedAt (api.raw = "updatedAt")
    17:  string title (api.raw = "title" )
    19:optional string type="" (api.raw = "type")
}
struct MenuResp{
   1:optional Menu data
   255:optional base.BaseResp baseResp
}
struct MenuListResp{
   1:optional list<Menu> data
   255:optional base.BaseResp baseResp
}
struct Role {
    1:optional i64 id,
    2:optional string name,
    3:optional string value,
    4:optional string defaultRouter,
    5:optional string remark,
    6:optional list<i64> apis,
}
struct RoleResp {
    1:optional Role data
    255:optional base.BaseResp baseResp
}
struct RoleListResp {
    1:optional list<Role> data
    255:optional base.BaseResp baseResp
}
struct CreateRoleReq{
    1:optional i64 id,
    2:optional string name,
    3:optional string value,
    4:optional string defaultRouter,
    5:optional string remark,
    6:optional list<i64> apis,
}

struct GetRoleListReq{
    1:optional base.PageReq pages
}
struct UpdateRoleReq {
    1:optional i64 id,
    2:optional string name,
    3:optional string value,
    4:optional string defaultRouter,
    5:optional string remark,
    6:optional list<i64> apis,
}
struct CreateMenuAuthReq{
    1: i64 roleId (api.raw = "roleId")
    2: list<i64> Ids (api.raw = "ids")
}
service SystemService  {
    ApiResp CreateApi(1: CreateApiReq req)
    ApiResp UpdateApi(1: UpdateApiReq req)
    base.NilResponse DeleteApi(1: base.IdReq req)
    ApiListResp ApiList(1: ApiListReq req)
    ApiListResp ApiTree(1: ApiListReq req)

    MenuResp CreateMenu(1: CreateMenuReq req)
    MenuResp UpdateMenu(1: UpdateMenuReq req)
    base.NilResponse DeleteMenu(1: base.IdReq req)
    MenuResp Menu(1: base.IdReq req)
    MenuListResp MenuList(1: MenuListReq req)
    MenuListResp MenuTree(1: MenuListReq req)

    RoleResp CreateRole(1: CreateRoleReq req)
    RoleResp GetRole(1: base.IdReq req)
    RoleResp LoginRole(1: base.CheckAccountReq req)
    RoleListResp GetRoleList(1: GetRoleListReq req)
    RoleResp UpdateRole(1: UpdateRoleReq req)
    base.NilResponse DeleteRole(1: base.IdReq req)
    // 创建菜单权限
    base.NilResponse CreateRoleMenu(1: CreateMenuAuthReq req)
    // 创建API权限
    base.NilResponse CreateRoleApi(1: CreateMenuAuthReq req)
    // 获取角色API列表
    MenuListResp GetRoleApi(1: base.IdReq req)
    //获取角色菜单列表
    MenuListResp GetRoleMenu(1: base.IdReq req)

}