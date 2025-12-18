namespace go crm
include "../base/base.thrift"
include "../base/crm.thrift"



//struct ClueResp{
//    1:optional Clue clue={},
//    255:optional base.BaseResp baseResp={}
//}
//struct CreateClueReq{
//
//}
//struct UpdateClueReq{
//    1:optional i64 id=0,
//}
//struct ClueListReq{
//    1:optional i64 page=1
//    2:optional i64 pageSize=10
//    3:optional string keyword=""
//}
//struct ClueListResp{
//    1:optional list<Clue> data= []
//    255:optional base.BaseResp baseResp={}
//}


struct FollowUpPlanResp{
    1:optional crm.FollowUpPlan data={},
    255:optional base.BaseResp baseResp={}
}
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
struct FollowUpPlanListResp{
    1:optional list<crm.FollowUpPlan> data= []
    255:optional base.BaseResp baseResp={}
}


struct FollowUpRecordResp{
    1:optional crm.FollowUpRecord data={},
    255:optional base.BaseResp baseResp={}
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
     256:optional i64 createdId=0,
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
     256:optional i64 createdId=0,
}
struct FollowUpRecordListReq{
    1:optional i64 page=1
    2:optional i64 pageSize=10
    3:optional string keyword=""
}
struct FollowUpRecordListResp{
    1:optional list<crm.FollowUpRecord> data= []
    255:optional base.BaseResp baseResp={}

}
struct OpportunitiesResp{
    1:optional crm.Opportunities data={},
    255:optional base.BaseResp baseResp={}
}
struct CreateOpportunitiesReq{
       2:optional i64 memberId=0,
       3:optional i64 userId=0,
       4:optional crm.OperatingPeriod period=0, //阶段
       7:optional string content="", // 内容
       8:optional i64 predictionAmount =0, //预测成交金额
       11:optional string title ="", //标题
       256:optional i64 createdId=0,
}
struct UpdateOpportunitiesReq{
    1:optional i64 id=0,
    2:optional i64 memberId=0,
    3:optional i64 userId=0,
    4:optional crm.OperatingPeriod period=0, //阶段
    7:optional string content="", // 内容
    8:optional i64 predictionAmount =0, //预测成交金额
    11:optional string title ="", //标题
    256:optional i64 createdId=0,
}
struct OpportunitiesListReq{
    1:optional i64 page=1
    2:optional i64 pageSize=10
    3:optional string keyword=""
}
struct OpportunitiesListResp{
    1:optional list<crm.Opportunities> data= []
    255:optional base.BaseResp baseResp={}
}

service CrmService  {
//      ClueResp GetClue(1:base.IdReq req)
//      ClueResp CreateClue(1: CreateClueReq req)
//      ClueResp UpdateClue(1: UpdateClueReq req)
//      base.BaseResp DeleteClue(1:base.IdReq req)
//      ClueListResp ClueList(1: ClueListReq req)


      FollowUpPlanResp GetFollowUpPlan(1:base.IdReq req)
      FollowUpPlanResp CreateFollowUpPlan(1: CreateFollowUpPlanReq req)
      FollowUpPlanResp UpdateFollowUpPlan(1: UpdateFollowUpPlanReq req)
      base.BaseResp DeleteFollowUpPlan(1:base.IdReq req)
      FollowUpPlanListResp FollowUpPlanList(1: FollowUpPlanListReq req)

      FollowUpRecordResp GetFollowUpRecord(1:base.IdReq req)
      FollowUpRecordResp CreateFollowUpRecord(1: CreateFollowUpRecordReq req)
      FollowUpRecordResp UpdateFollowUpRecord(1: UpdateFollowUpRecordReq req)
      base.BaseResp DeleteFollowUpRecord(1:base.IdReq req)
      FollowUpRecordListResp FollowUpRecordList(1:  FollowUpRecordListReq req)

      OpportunitiesResp GetOpportunities(1:base.IdReq req)
      OpportunitiesResp CreateOpportunities(1: CreateOpportunitiesReq req)
      OpportunitiesResp UpdateOpportunities(1: UpdateOpportunitiesReq req)
      base.BaseResp DeleteOpportunities(1:base.IdReq req)
      OpportunitiesListResp OpportunitiesList(1: OpportunitiesListReq req)

}