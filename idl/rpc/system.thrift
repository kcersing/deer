namespace go System
include "../base/base.thrift"

struct Api{

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

}
struct CreateApiReq{

}
struct UpdateApiReq{

}
struct CreateMenuReq{

}
struct UpdateMenuReq{

}
struct MenuListReq{

}
struct MenuResp{

}
struct MenuListResp{

}

service SystemService  {
      ApiResp CreateApi(1: CreateApiReq req)
      ApiResp UpdateApi(1: UpdateApiReq req)
      base.BaseResp DeleteApi(1: base.IdReq req)
      ApiListResp ApiList(1: ApiListReq req)
      ApiListResp ApiTree(1: ApiListReq req)

      MenuResp CreateMenu(1: CreateMenuReq req)
      MenuResp UpdateMenu(1: UpdateMenuReq req)
      base.BaseResp DeleteMenu(1: base.IdReq req)
      MenuResp Menu(1: base.IdReq req)
      MenuListResp MenuList(1: MenuListReq req)
      MenuListResp MenuTree(1: MenuListReq req)



}