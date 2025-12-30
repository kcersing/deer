namespace go schedule

include "../base/base.thrift"
include "../base/schedule.thrift"

struct CreateScheduleReq{

    2:optional string type ="" (api.raw = "type")
    10:optional i64 userId =0 (api.raw = "userId")
    3:optional i64 productItemId =0 (api.raw = "productItemId")
    4:optional i64 venueId =0 (api.raw = "venueId")
    5:optional i64 placeId  =0(api.raw = "placeId")
    6:optional i64 num =0 (api.raw = "num")
    7:optional string startAt="" (api.raw = "startAt")
    8:optional double price =0 (api.raw = "price")
    14:optional i64 status = 0
    11:optional i64 memberId =0 (api.raw = "memberId")
    12:optional i64 memberProductId =0 (api.raw = "memberProductId")
    13:optional i64 memberProductItemId =0 (api.raw = "memberProductItemId")
}
struct UpdateScheduleReq{

}

struct GetScheduleListReq{
     1:optional i64 page=1
     2:optional i64 pageSize=10

}
struct GetScheduleDateListReq{
     1:optional i64 page=1
     2:optional i64 pageSize=10

}
struct GetMemberScheduleListReq{
     1:optional i64 page=1
     2:optional i64 pageSize=10

}
struct GetUserScheduleListReq{
     1:optional i64 page=1
     2:optional i64 pageSize=10

}
struct ScheduleResp{
    1:optional schedule.Schedule data= {}
    255:optional base.BaseResp baseResp={}
}
struct ScheduleListResp{
     1:optional list<schedule.Schedule> data= []
     255:optional base.BaseResp baseResp={}
}
struct ScheduleDateListResp{
     1:optional list<schedule.Schedule> data= []
    255:optional base.BaseResp baseResp={}
}
struct MemberScheduleResp{
    1:optional schedule.MemberSchedule data= {}
    255:optional base.BaseResp baseResp={}
}
struct MemberScheduleListResp{
     1:optional list<schedule.MemberSchedule> data= []
    255:optional base.BaseResp baseResp={}
}
struct UserScheduleResp{
    1:optional schedule.UserSchedule data= {}
    255:optional base.BaseResp baseResp={}
}
struct UserScheduleListResp{
    1:optional list<schedule.UserSchedule> data= []
    255:optional base.BaseResp baseResp={}
}


service ScheduleService {

    ScheduleResp CreateSchedule(1: CreateScheduleReq req)
    ScheduleResp UpdateSchedule(1: UpdateScheduleReq req)
    base.NilResponse DeleteSchedule(1: base.IdReq req)
    ScheduleResp GetSchedule(1: base.IdReq req)
    ScheduleListResp GetScheduleList(1: GetScheduleListReq req)
    ScheduleDateListResp GetScheduleDateList(1: GetScheduleDateListReq req )

    MemberScheduleListResp GetMemberScheduleList(1: GetMemberScheduleListReq req)
    UserScheduleListResp GetUserScheduleList(1: GetUserScheduleListReq req)

    /**取消课程列表*/
    base.NilResponse CancelSchedule(1: base.IdReq req)
    /**取消会员课程*/
    base.NilResponse CancelMemberSchedule(1: base.IdReq req)
    /**会员签到*/
    base.NilResponse SignMemberSchedule(1: base.IdReq req)
    /**教练签到*/
    base.NilResponse SignUserSchedule(1: base.IdReq req)


}