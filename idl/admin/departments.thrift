namespace go departments
include "../base/base.thrift"


struct CreateDepartmentsReq{
     2:optional string name=""
     3:optional string managerId=""
     4:optional string parentId=""
     5:optional string desc=""
     7:optional i64 status=0
}
struct UpdateDepartmentsReq{
     1:optional i64 id=0
     2:optional string name=""
     3:optional string managerId=""
     4:optional string parentId=""
     5:optional string desc=""
     7:optional i64 status=0
}
struct GetDepartmentsListReq{
      1:optional i64 page=1
      2:optional i64 pageSize=10
      3:optional string keyword=""
}
service UserService  {
    base.NilResponse CreateDepartments(1: CreateDepartmentsReq req)(api.post = "/service/departments/create")
    base.NilResponse DeleteDepartments(1: base.IdReq req)(api.post = "/service/departments/delete")
    base.NilResponse UpdateDepartments(1: UpdateDepartmentsReq req)(api.post = "/service/departments/update")
    base.NilResponse GetDepartments(1: base.IdReq req)(api.post = "/service/departments")
    base.NilResponse GetDepartmentsList(1: GetDepartmentsListReq req)(api.post = "/service/departments/list")
}