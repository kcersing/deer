import {Position} from "@/pages/opm/positions/service/data";
import {BaseResp} from "@/services/typings";

export type Departments = {
  id?: number;
  name?: string;
  managerId?:number;
  parentId?: number;
  desc?: string;
  status?: number;
  createdAt?: string;
  updatedAt?: string;
  createdId?: number;
};
export type DepartmentsResp = {
  data?: Departments;
  baseResp?: BaseResp;
}
export type DepartmentsListResp={
  data?: Departments[];
  baseResp?: BaseResp;
}

export type CreateDepartmentsReq = {
  name?: string;
  managerId?:number;
  parentId?: number;
  desc?: string;
  status?: number;
};

export type UpdateDepartmentsReq = {
  id?: number;
  name?: string;
  managerId?:number;
  parentId?: number;
  desc?: string;
  status?: number;
};
