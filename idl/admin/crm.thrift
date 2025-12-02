namespace go crm
include "../base/base.thrift"
include "../base/crm.thrift"


struct CreateFollowUpPlanReq{
    2:optional string content="",//计划内容
    3:optional string time="", //计划时间
    4:optional i64 memberId=0, //跟进客户
    6:optional i64 userId=0,// 执行人
    7:optional i64 status=0, //计划状态 待完成 延期 完成
    8:optional i64 createdId=0, //计划制定人
    11:optional i64 division=0, //部门
}
struct UpdateFollowUpPlanReq{
    1:optional i64 id=0,
    2:optional string content="",//计划内容
    3:optional string time="", //计划时间
    4:optional i64 memberId=0, //跟进客户
    6:optional i64 userId=0,// 执行人
    7:optional i64 status=0, //计划状态 待完成 延期 完成
    8:optional i64 createdId=0, //计划制定人
    11:optional i64 division=0, //部门
}
struct FollowUpPlanListReq{
    1:optional i64 page=1
    2:optional i64 pageSize=10
    3:optional string keyword=""
}

struct CreateFollowUpRecordReq{
     2:optional string content="",
     3:optional i64 followUpId=0,
     4:optional i64 method=0,
     5:optional i64 status=0,
     6:optional i64 userId=0,
     7:optional i64 division=0,
     10:optional string record ="",
     11:optional i64 opportunitiesId=0,
}
struct UpdateFollowUpRecordReq{
    1:optional i64 id=0,
     2:optional string content="",
     3:optional i64 followUpId=0,
     4:optional i64 method=0,
     5:optional i64 status=0,
     6:optional i64 userId=0,
     7:optional i64 division=0,
     10:optional string record ="",
     11:optional i64 opportunitiesId=0,
}
struct FollowUpRecordListReq{
    1:optional i64 page=1
    2:optional i64 pageSize=10
    3:optional string keyword=""
}
struct CreateOpportunitiesReq{
       2:optional i64 memberId=0,
       3:optional i64 userId=0,
       4:optional crm.OperatingPeriod period=0, //阶段
       7:optional string content="", // 内容
       8:optional i64 predictionAmount =0, //预测成交金额
       11:optional string title ="", //标题
}
struct UpdateOpportunitiesReq{
    1:optional i64 id=0,
    2:optional i64 memberId=0,
    3:optional i64 userId=0,
    4:optional crm.OperatingPeriod period=0, //阶段
    7:optional string content="", // 内容
    8:optional i64 predictionAmount =0, //预测成交金额
    11:optional string title ="", //标题
}
struct OpportunitiesListReq{
    1:optional i64 page=1
    2:optional i64 pageSize=10
    3:optional string keyword=""
}
service CrmService {
     base.NilResponse GetFollowUpPlan(1:base.IdReq req)
     base.NilResponse CreateFollowUpPlan(1: CreateFollowUpPlanReq req)
     base.NilResponse UpdateFollowUpPlan(1: UpdateFollowUpPlanReq req)
     base.NilResponse DeleteFollowUpPlan(1:base.IdReq req)
     base.NilResponse FollowUpPlanList(1: FollowUpPlanListReq req)

     base.NilResponse GetFollowUpRecord(1:base.IdReq req)
     base.NilResponse CreateFollowUpRecord(1: CreateFollowUpRecordReq req)
     base.NilResponse UpdateFollowUpRecord(1: UpdateFollowUpRecordReq req)
     base.NilResponse DeleteFollowUpRecord(1:base.IdReq req)
     base.NilResponse FollowUpRecordList(1:  FollowUpRecordListReq req)

     base.NilResponse GetOpportunities(1:base.IdReq req)
     base.NilResponse CreateOpportunities(1: CreateOpportunitiesReq req)
     base.NilResponse UpdateOpportunities(1: UpdateOpportunitiesReq req)
     base.NilResponse DeleteOpportunities(1:base.IdReq req)
     base.NilResponse OpportunitiesList(1: OpportunitiesListReq req)
 }