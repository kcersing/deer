package main

import (
	"context"
	base "gen/kitex_gen/base"
	crm "gen/kitex_gen/crm"

	"crm/biz/service"
)

// CrmServiceImpl implements the last service interface defined in the IDL.
type CrmServiceImpl struct{}

// GetFollowUpPlan implements the CrmServiceImpl interface.
func (s *CrmServiceImpl) GetFollowUpPlan(ctx context.Context, req *base.IdReq) (resp *crm.FollowUpPlanResp, err error) {
	resp, err = service.NewGetFollowUpPlanService(ctx).Run(req)

	return resp, err
}

// CreateFollowUpPlan implements the CrmServiceImpl interface.
func (s *CrmServiceImpl) CreateFollowUpPlan(ctx context.Context, req *crm.CreateFollowUpPlanReq) (resp *crm.FollowUpPlanResp, err error) {
	resp, err = service.NewCreateFollowUpPlanService(ctx).Run(req)

	return resp, err
}

// UpdateFollowUpPlan implements the CrmServiceImpl interface.
func (s *CrmServiceImpl) UpdateFollowUpPlan(ctx context.Context, req *crm.UpdateFollowUpPlanReq) (resp *crm.FollowUpPlanResp, err error) {
	resp, err = service.NewUpdateFollowUpPlanService(ctx).Run(req)

	return resp, err
}

// DeleteFollowUpPlan implements the CrmServiceImpl interface.
func (s *CrmServiceImpl) DeleteFollowUpPlan(ctx context.Context, req *base.IdReq) (resp *base.BaseResp, err error) {
	resp, err = service.NewDeleteFollowUpPlanService(ctx).Run(req)

	return resp, err
}

// FollowUpPlanList implements the CrmServiceImpl interface.
func (s *CrmServiceImpl) FollowUpPlanList(ctx context.Context, req *crm.FollowUpPlanListReq) (resp *crm.FollowUpPlanListResp, err error) {
	resp, err = service.NewFollowUpPlanListService(ctx).Run(req)

	return resp, err
}

// GetFollowUpRecord implements the CrmServiceImpl interface.
func (s *CrmServiceImpl) GetFollowUpRecord(ctx context.Context, req *base.IdReq) (resp *crm.FollowUpRecordResp, err error) {
	resp, err = service.NewGetFollowUpRecordService(ctx).Run(req)

	return resp, err
}

// CreateFollowUpRecord implements the CrmServiceImpl interface.
func (s *CrmServiceImpl) CreateFollowUpRecord(ctx context.Context, req *crm.CreateFollowUpRecordReq) (resp *crm.FollowUpRecordResp, err error) {
	resp, err = service.NewCreateFollowUpRecordService(ctx).Run(req)

	return resp, err
}

// UpdateFollowUpRecord implements the CrmServiceImpl interface.
func (s *CrmServiceImpl) UpdateFollowUpRecord(ctx context.Context, req *crm.UpdateFollowUpRecordReq) (resp *crm.FollowUpRecordResp, err error) {
	resp, err = service.NewUpdateFollowUpRecordService(ctx).Run(req)

	return resp, err
}

// DeleteFollowUpRecord implements the CrmServiceImpl interface.
func (s *CrmServiceImpl) DeleteFollowUpRecord(ctx context.Context, req *base.IdReq) (resp *base.BaseResp, err error) {
	resp, err = service.NewDeleteFollowUpRecordService(ctx).Run(req)

	return resp, err
}

// FollowUpRecordList implements the CrmServiceImpl interface.
func (s *CrmServiceImpl) FollowUpRecordList(ctx context.Context, req *crm.FollowUpRecordListReq) (resp *crm.FollowUpRecordListResp, err error) {
	resp, err = service.NewFollowUpRecordListService(ctx).Run(req)

	return resp, err
}

// GetOpportunities implements the CrmServiceImpl interface.
func (s *CrmServiceImpl) GetOpportunities(ctx context.Context, req *base.IdReq) (resp *crm.OpportunitiesResp, err error) {
	resp, err = service.NewGetOpportunitiesService(ctx).Run(req)

	return resp, err
}

// CreateOpportunities implements the CrmServiceImpl interface.
func (s *CrmServiceImpl) CreateOpportunities(ctx context.Context, req *crm.CreateOpportunitiesReq) (resp *crm.OpportunitiesResp, err error) {
	resp, err = service.NewCreateOpportunitiesService(ctx).Run(req)

	return resp, err
}

// UpdateOpportunities implements the CrmServiceImpl interface.
func (s *CrmServiceImpl) UpdateOpportunities(ctx context.Context, req *crm.UpdateOpportunitiesReq) (resp *crm.OpportunitiesResp, err error) {
	resp, err = service.NewUpdateOpportunitiesService(ctx).Run(req)

	return resp, err
}

// DeleteOpportunities implements the CrmServiceImpl interface.
func (s *CrmServiceImpl) DeleteOpportunities(ctx context.Context, req *base.IdReq) (resp *base.BaseResp, err error) {
	resp, err = service.NewDeleteOpportunitiesService(ctx).Run(req)

	return resp, err
}

// OpportunitiesList implements the CrmServiceImpl interface.
func (s *CrmServiceImpl) OpportunitiesList(ctx context.Context, req *crm.OpportunitiesListReq) (resp *crm.OpportunitiesListResp, err error) {
	resp, err = service.NewOpportunitiesListService(ctx).Run(req)

	return resp, err
}
