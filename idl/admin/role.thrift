namespace go role
include "../base/base.thrift"
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
service  RoleService  {
    base.NilResponse CreateRole(1: CreateRoleReq req)(api.post = "/service/role/create")
    base.NilResponse GetRole(1: base.IdReq req)(api.post = "/service/role")
    base.NilResponse GetRoleList(1: GetRoleListReq req)(api.post = "/service/role/list")
    base.NilResponse UpdateRole(1: UpdateRoleReq req)(api.post = "/service/role/update")

    base.NilResponse DeleteRole(1: base.IdReq req) (api.post = "/service/role/delete")
    // 创建菜单权限
    base.NilResponse CreateRoleMenu(1: CreateMenuAuthReq req)(api.post = "/service/role/create/menu")
    // 创建API权限
    base.NilResponse CreateRoleApi(1: CreateMenuAuthReq req)(api.post = "/service/role/create/api")
    // 获取角色API列表
    base.NilResponse GetRoleApi(1: base.IdReq req)(api.post = "/service/role/api")
    //获取角色菜单列表
    base.NilResponse GetRoleMenu(1: base.IdReq req)(api.post = "/service/role/menu")
}