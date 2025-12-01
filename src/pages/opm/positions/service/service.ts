import { request } from '@umijs/max';
import {PositionsListResp,CreatePositionsReq,PositionsResp,UpdatePositionsReq} from "./data";
import {BaseResp} from "@/services/typings";
import {headers} from "@/services/ant-design-pro/utils";

/** 创建 POST /service/positions/create */
export async function createPositions(body: CreatePositionsReq, options?: { [key: string]: any }) {
  return request<PositionsResp>('/service/positions/create', {
    method: 'POST',
    headers: {
      ...headers,
    },
    data: body,
    ...(options || {}),
  });
}
/** 获取信息 POST /service/positions */
export async function getPositions(body: { id:number }, options?: { [key: string]: any }) {
  return request<PositionsResp>('/service/positions', {
    method: 'POST',
    headers: {
      ...headers,
    },
    data: body,
    ...(options || {}),
  });
}
/** 获取列表 POST /service/positions/list*/
export async function getPositionsList(
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
  return request<PositionsListResp>('/service/positions/list', {
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

/** 更新role  POST /service/positions/update*/
export async function updatePositions(body: UpdatePositionsReq, options?: { [key: string]: any }) {
  return request<PositionsResp>('/service/positions/update', {
    method: 'POST', headers: {
      ...headers,
    }, ...(options || {}),
  });
}

/** 删除role  POST /service/positions/delete*/
export async function deletePositions(options?: { [key: string]: any }) {
  return request<BaseResp>('/service/positions/delete', {
    method: 'POST', headers: {
      ...headers,
    }, ...(options || {}),
  });
}
