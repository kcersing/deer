import { request } from '@umijs/max';
import {DepartmentsListResp,CreateDepartmentsReq,DepartmentsResp,UpdateDepartmentsReq} from "./data";
import {BaseResp} from "@/services/typings";
import {headers} from "@/services/ant-design-pro/utils";

/** 创建 POST /service/departments/create */
export async function createDepartments(body: CreateDepartmentsReq, options?: { [key: string]: any }) {
  return request<DepartmentsResp>('/service/departments/create', {
    method: 'POST',
    headers: {
      ...headers,
    },
    data: body,
    ...(options || {}),
  });
}
/** 获取信息 POST /service/departments */
export async function getDepartments(body: { id:number }, options?: { [key: string]: any }) {
  return request<DepartmentsResp>('/service/departments', {
    method: 'POST',
    headers: {
      ...headers,
    },
    data: body,
    ...(options || {}),
  });
}
/** 获取列表 POST /service/departments/list*/
export async function getDepartmentsList(
  params: {
    // query
    /** 当前的页码 */
    current?: number;
    /** 页面的容量 */
    pageSize?: number;
    keywords?: string;

  },
  options?: { [key: string]: any },
) {
  return request<DepartmentsListResp>('/service/departments/list', {
    method: 'POST',
    headers: {
      ...headers,
    },
    data: {
      page: params.current,
      ...params,
    },
    ...(options || {}),
  });
}

/** 更新  POST /service/departments/update*/
export async function updateDepartments(body: UpdateDepartmentsReq, options?: { [key: string]: any }) {
  return request<DepartmentsResp>('/service/departments/update', {
    method: 'POST', headers: {
      ...headers,
    }, ...(options || {}),
  });
}

/** 删除  POST /service/departments/delete*/
export async function deleteDepartments(options?: { [key: string]: any }) {
  return request<BaseResp>('/service/departments/delete', {
    method: 'POST', headers: {
      ...headers,
    }, ...(options || {}),
  });
}
