namespace go Role
include "../base/base.thrift"
include "system.thrift"

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
struct MenuAuthReq{
    1: i64 roleId (api.raw = "roleId")
    2: list<i64> Ids (api.raw = "ids")
}
service RoleService  {
    RoleResp CreateRole(1: CreateRoleReq req)
    RoleResp GetRole(1: base.IdReq req)
    RoleResp LoginRole(1: base.CheckAccountReq req)
    RoleListResp GetRoleList(1: GetRoleListReq req)
    RoleResp UpdateRole(1: UpdateRoleReq req)
    base.BaseResp DeleteRole(1: base.IdReq req)
    // 创建菜单权限
    base.BaseResp CreateMenu(1: MenuAuthReq req)
    // 创建API权限
    base.BaseResp CreateApi(1: MenuAuthReq req)
    // 获取角色API列表
    system.MenuListResp GetApi(1: base.IdReq req)
    //获取角色菜单列表
   system.MenuListResp GetMenu(1: base.IdReq req)


}
