namespace go user
include "../base/base.thrift"
include "../base/user.thrift"


struct UserResp {
    1:optional user.User data={}
    255:optional base.BaseResp baseResp={}
}
struct UserListResp {
    1:optional list<user.User> data= []
    255:optional base.BaseResp baseResp={}
}
struct CreateUserReq{
     1: string username,
     2:optional string avatar="",
     3:optional string mobile="",
     4:optional string name="",
     5:optional i64 status=0,
     6:optional i64 gender=0,
     7:optional string birthday="",
     8:optional string  desc="",

     9:optional i64 departmentsId=0,
     10:optional i64 positionsId=0,
     256:optional i64 createdId=0,
}

struct GetUserListReq{
      1:optional i64 page=1
      2:optional i64 pageSize=10
      3:optional string keyword=""
      4:optional string name=""
      5:optional string mobile="",
      6:optional list<i64> tags=[]
}
struct UpdateUserReq {
    1:optional i64 id=0,
    2:optional string avatar="",
    3:optional string mobile="",
    4:optional string name="",
    5:optional i64 status=0,
    6:optional i64 gender=0,
    7:optional string birthday="",
    8:optional string   desc="",

    9:optional i64 departmentsId=0,
    10:optional i64 positionsId=0,
     256:optional i64 createdId=0,

}

struct ChangePasswordReq {
    1:optional i64 id=0,
    2:optional string password="",
}
struct SetUserRoleReq{
    1:optional i64 id=0,
    2:optional i64 roleId=0,
     256:optional i64 createdId=0,

}


struct DepartmentsResp {
    1:optional user.Departments data={}
    255:optional base.BaseResp baseResp={}
}
struct DepartmentsListResp {
    1:optional list<user.Departments> data= []
    255:optional base.BaseResp baseResp={}
}

struct CreateDepartmentsReq{
     2:optional string name="",
     3:optional i64 managerId=0,
     4:optional i64 parentId=0,
     5:optional string desc="",
     7:optional i64 status=0,
     256:optional i64 createdId=0,

}
struct UpdateDepartmentsReq{
     1:optional i64 id=0,
     2:optional string name="",
     3:optional i64 managerId=0,
     4:optional i64 parentId=0,
     5:optional string desc="",
     7:optional i64 status=0,
     256:optional i64 createdId=0,

}
struct GetDepartmentsListReq{
      1:optional i64 page=1
      2:optional i64 pageSize=10
      3:optional string keyword=""
}
struct CreatePositionsReq{
     2:optional string name="",
     3:optional string code="",
     4:optional i64 departmentId=0,
     5:optional i64 parentId=0,
     6:optional string desc="",
     7:optional i64 status=0,
     9:optional i64 quota=0,
     256:optional i64 createdId=0,

}
struct UpdatePositionsReq{
     1:optional i64 id=0,
     2:optional string name="",
     3:optional string code="",
     4:optional i64 departmentId=0,
     5:optional i64 parentId=0,
     6:optional string desc="",
     7:optional i64 status=0,
     9:optional i64 quota=0,
     256:optional i64 createdId=0,
}
struct GetPositionsListReq{
      1:optional i64 page=1
      2:optional i64 pageSize=10
      3:optional string keyword=""
}

struct PositionsResp {
    1:optional user.Positions data={}
    255:optional base.BaseResp baseResp={}
}
struct PositionsListResp {
    1:optional list<user.Positions> data= []
    255:optional base.BaseResp baseResp={}
}
struct UserIdsResp{
    1:optional list<i64> data= []
    255:optional base.BaseResp baseResp={}
}

service UserService  {
    UserResp CreateUser(1: CreateUserReq req)
    UserResp GetUser(1: base.IdReq req)
    UserResp LoginUser(1: base.CheckAccountReq req)
    UserListResp GetUserList(1: GetUserListReq req)
    UserResp UpdateUser(1: UpdateUserReq req)
    
    base.NilResponse ChangePassword(1: ChangePasswordReq req)
    base.NilResponse DeleteUser(1: base.IdReq req)
    base.NilResponse SetUserRole(1: SetUserRoleReq  req)


    DepartmentsResp CreateDepartments(1: CreateDepartmentsReq req)
    base.NilResponse DeleteDepartments(1: base.IdReq req)
    DepartmentsResp UpdateDepartments(1: UpdateDepartmentsReq req)
    DepartmentsResp GetDepartments(1: base.IdReq req)
    DepartmentsListResp GetDepartmentsList(1: GetDepartmentsListReq req)


    PositionsResp CreatePositions(1: CreatePositionsReq req)
    base.NilResponse DeletePositions(1: base.IdReq req)
    PositionsResp UpdatePositions(1: UpdatePositionsReq req)
    PositionsResp GetPositions(1: base.IdReq req)
    PositionsListResp GetPositionsList(1: GetPositionsListReq req)


    UserIdsResp GetUserIds(1: GetUserListReq req)
}
