namespace go menu

include "../base/base.thrift"

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

    12:optional i64 status=0 (api.raw = "status")
    13:optional string icon="" (api.raw = "icon")




}
struct UpdateMenuReq{
    1: optional i64 id=0 (api.raw = "id")
    2:optional string name="" (api.raw = "name" api.vd = "len($) > 0 && len($) < 33>")
    3:optional i64 parentId=0 (api.raw = "parentId")
    4:optional i64 level=0 (api.raw = "level")
    5:optional string path="" (api.raw = "path")
    6:optional string redirect="" (api.raw = "redirect")
    7:optional string component="" (api.raw = "component")
    8:optional i64 menuType=0 (api.raw = "menuType")
    9:optional i64 orderNo=0 (api.raw = "orderNo")
    10:optional i64 ignore=0 (api.raw = "ignore")

    12:optional i64 status=0 (api.raw = "status")
    13:optional string icon="" (api.raw = "icon")
}
struct MenuListReq{
    1:  optional i64 page=1 (api.raw = "page")
    2:  optional i64 pageSize=100 (api.raw = "pageSize")
}

service MenuService{
    base.NilResponse CreateMenu(1: CreateMenuReq req)(api.post = "/service/menu/create")
    base.NilResponse UpdateMenu(1: UpdateMenuReq req)(api.post = "/service/menu/update")
    base.NilResponse DeleteMenu(1: base.IdReq req)(api.post = "/service/menu/delete")
    base.NilResponse GetMenu(1: base.IdReq req)(api.post = "/service/menu")
    base.NilResponse MenuList(1: MenuListReq req)(api.post = "/service/menu/list")
    base.NilResponse MenuTree(1: MenuListReq req) (api.post = "/service/menu/tree")
}