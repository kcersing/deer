namespace go base
include "system.thrift"
struct User {
     1:optional i64 id=0,
     2:optional string username="",
     3:optional string password="",
     4:optional string avatar="",
     5:optional string mobile="",
     6:optional string name="",
     7:optional i64 status=0,
     9:optional i64 gender=0,
     10:optional string birthday="",

     11:optional string  lastAt="",//最后一次登录时间
     12:optional string  lastIp="",//最后一次登录ip

     13:optional string  desc="",//详情
     14:optional list<system.Role> roles= [],//角色
     15:optional i64 departmentsId=0,
     16:optional i64 positionsId=0,

     251:optional string createdAt="",
     252:optional string updatedAt="",
     253:optional i64 createdId=0,
}

struct Departments {
     1:optional i64 id=0,
     2:optional string name="",
     3:optional i64 managerId=0,
     4:optional i64 parentId=0,
     5:optional string desc="",
     7:optional i64 status=0,
     251:optional string createdAt="",
     252:optional string updatedAt="",
     253:optional i64 createdId=0,
}
struct Positions {
     1:optional i64 id=0,
     2:optional string name="",
     3:optional string code="",
     4:optional i64 departmentId=0,
     5:optional i64 parentId=0,
     6:optional string desc="",
     7:optional i64 status=0,
     9:optional i64 quota=0,

     251:optional string createdAt="",
     252:optional string updatedAt="",
     253:optional i64 createdId=0,
}

