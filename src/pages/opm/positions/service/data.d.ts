export type UpdatePositionsReq = {
}

export type Positions = {
  id?: number;
  name?: string;
  code?: string;
  departmentId?: number;
  parentId?: number;
  desc?: string;
  status?: number;
  quota?: number;
  createdAt?: string;
  updatedAt?: string;
  createdId?: number;
}
export type CreatePositionsReq = {
  name?: string;
  code?: string;
  departmentId?: number;
  parentId?: number;
  desc?: string;
  status?: number;
  quota?: number;
}
export type PositionsResp = {
  data?: Positions;
  baseResp?: BaseResp;
}
export type PositionsListResp={
  data?: Position[];
  baseResp?: BaseResp;
}
