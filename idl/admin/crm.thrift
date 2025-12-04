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
     base.NilResponse GetFollowUpPlan(1:base.IdReq req)(api.post = "/service/follow-up-plan")
     base.NilResponse CreateFollowUpPlan(1: CreateFollowUpPlanReq req)(api.post = "/service/follow-up-plan/create")
     base.NilResponse UpdateFollowUpPlan(1: UpdateFollowUpPlanReq req)(api.post = "/service/follow-up-plan/update")
     base.NilResponse DeleteFollowUpPlan(1:base.IdReq req)(api.post = "/service/follow-up-plan/delete")
     base.NilResponse FollowUpPlanList(1: FollowUpPlanListReq req)(api.post = "/service/follow-up-plan/list")

     base.NilResponse GetFollowUpRecord(1:base.IdReq req)(api.post = "/service/follow-up-record")
     base.NilResponse CreateFollowUpRecord(1: CreateFollowUpRecordReq req)(api.post = "/service/follow-up-record/create")
     base.NilResponse UpdateFollowUpRecord(1: UpdateFollowUpRecordReq req)(api.post = "/service/follow-up-record/update")
     base.NilResponse DeleteFollowUpRecord(1:base.IdReq req)(api.post = "/service/follow-up-record/delete")
     base.NilResponse FollowUpRecordList(1:  FollowUpRecordListReq req)(api.post = "/service/follow-up-record/list")

     base.NilResponse GetOpportunities(1:base.IdReq req)(api.post = "/service/opportunities")
     base.NilResponse CreateOpportunities(1: CreateOpportunitiesReq req)(api.post = "/service/opportunities/create")
     base.NilResponse UpdateOpportunities(1: UpdateOpportunitiesReq req)(api.post = "/service/opportunities/update")
     base.NilResponse DeleteOpportunities(1:base.IdReq req)(api.post = "/service/opportunities/delete")
     base.NilResponse OpportunitiesList(1: OpportunitiesListReq req)(api.post = "/service/opportunities/list")
 }