namespace go crm
include "../base/base.thrift"


//线索
//struct Clue {
//    1:optional i64 id=0,
//}

//跟进计划
struct FollowUpPlan {
    1:optional i64 id=0,
    2:optional string content="",//计划内容
    3:optional string time="", //计划时间
    4:optional i64 memberId=0, //跟进客户
    6:optional i64 userId=0,// 执行人
    7:optional i64 status=0, //计划状态 待完成 延期 完成
    8:optional i64 createdId=0, //计划制定人
    9:optional i64 createdAt=0, //创建时间
    10:optional i64 updatedAt=0, //更新时间
    11:optional i64 division=0, //部门
}

//跟进记录
struct GenFollowUpRecord {
     1:optional i64 id=0,
     2:optional i64 content="", //跟进记录内容 线索筛选，需求挖掘、跟进方案制定、成交转化、后续服务
     3:optional i64 followUpId=0, //跟进计划
     4:optional i64 method=0, //跟进方式
     5:optional i64 status=0, //跟进状态
     6:optional i64 userId=0, //跟进人
     7:optional i64 division=0, //部门
     8:optional i64 createdAt=0, //创建时间
     9:optional i64 updatedAt=0, //更新时间
     10:optional string record ="", //跟进记录
     11:optional i64 opportunitiesId=0, //跟进商机
     // 12:optional i64 AAA=0, //跟进主体
}
//商机
struct Opportunities {
     1:optional i64 id=0,
     2:optional i64 memberId=0,
     3:optional i64 userId=0,
     4:optional OperatingPeriod period=0, //阶段
     5:optional string  periodTime="", //阶段变更时间
     6:optional i64 winRate=0, //赢率
     7:optional i64 content="", // 内容
     8:optional i64 predictionAmount =0, //预测成交金额
     9:optional i64 createdAt=0, //创建时间
     10:optional i64 updatedAt=0, //更新时间
     11:optional string title ="", //标题

//      预计成交日期
//      协作人
//      输单原因
//      商品明细



}
enum OperatingPeriod {
    IdentifyCustomerNeeds         = 0   //需求发现‌  30%
    ValidateCustomerRequirements  = 1   //‌需求确认‌  40%
    ProposalAndQuotation          = 2   //方案报价‌  60%
    BusinessNegotiation           = 3   //商务谈判‌  80%
    Invalid                       = 4   //无效 0%
    LoseTheDeal                   = 5   //输单‌ 0%
    WinTheDeal                    = 6   //赢单‌   100%
}
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
    1:optional FollowUpPlan followUpPlan={},
    255:optional base.BaseResp baseResp={}
}
struct CreateFollowUpPlanReq{

}
struct UpdateFollowUpPlanReq{
    1:optional i64 id=0,
}
struct FollowUpPlanListReq{
    1:optional i64 page=1
    2:optional i64 pageSize=10
    3:optional string keyword=""
}
struct FollowUpPlanListResp{
    1:optional list<FollowUpPlan> data= []
    255:optional base.BaseResp baseResp={}
}


struct GenFollowUpRecordResp{
    1:optional GenFollowUpRecord followUpRecord={},
    255:optional base.BaseResp baseResp={}
}
struct CreateGenFollowUpRecordReq{

}
struct UpdateGenFollowUpRecordReq{
    1:optional i64 id=0,
}
struct GenFollowUpRecordListReq{
    1:optional i64 page=1
    2:optional i64 pageSize=10
    3:optional string keyword=""
}
struct GenFollowUpRecordListResp{
    1:optional list<GenFollowUpRecord> data= []
    255:optional base.BaseResp baseResp={}

}
struct OpportunitiesResp{
    1:optional Opportunities opportunities={},
    255:optional base.BaseResp baseResp={}
}
struct CreateOpportunitiesReq{
 1:optional i64 memberId=0,
 2:optional i64 userId=0,
 3:optional OperatingPeriod period=0,
// 4:optional i64 period =0,
 5:optional i64 winRate=0,
}
struct UpdateOpportunitiesReq{
    1:optional i64 id=0,
}
struct OpportunitiesListReq{
    1:optional i64 page=1
    2:optional i64 pageSize=10
    3:optional string keyword=""
}
struct OpportunitiesListResp{
    1:optional list<Opportunities> data= []
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

      GenFollowUpRecordResp GetGenFollowUpRecord(1:base.IdReq req)
      GenFollowUpRecordResp CreateGenFollowUpRecord(1: CreateGenFollowUpRecordReq req)
      GenFollowUpRecordResp UpdateGenFollowUpRecord(1: UpdateGenFollowUpRecordReq req)
      base.BaseResp DeleteGenFollowUpRecord(1:base.IdReq req)
      GenFollowUpRecordListResp GenFollowUpRecordList(1:  GenFollowUpRecordListReq req)

      OpportunitiesResp GetOpportunities(1:base.IdReq req)
      OpportunitiesResp CreateOpportunities(1: CreateOpportunitiesReq req)
      OpportunitiesResp UpdateOpportunities(1: UpdateOpportunitiesReq req)
      base.BaseResp DeleteOpportunities(1:base.IdReq req)
      OpportunitiesListResp OpportunitiesList(1: OpportunitiesListReq req)

}