import {request} from '@umijs/max';
import {ApiListResp,Api,ApiResp} from "./data";
import {BaseResp} from "@/services/typings";
import {headers} from "@/services/ant-design-pro/utils";


/** 更新Api  POST /service/api/update*/
export async function updateApi(options?: { [key: string]: any }) {
  return request<ApiResp>('/service/api/update', {
    method: 'POST', headers: {
      ...headers,
    },
    data: {
      ...options,
    },
    ...(options || {}),
  });
}

/** 创建 Api  POST /service/api/create*/
export async function createApi(options?: { [key: string]: any }) {
  return request<ApiResp>('/service/api/create', {
    method: 'POST', headers: {
      ...headers,
    },
    data: {
      ...options,
    },
    ...(options || {}),
  });
}


/** 获取api 列表 POST /service/api/list*/
export async function getApiList(params: {
  // query
  /** 当前的页码 */
  current?: number; /** 页面的容量 */
  pageSize?: number; keywords?: string;

}, options?: { [key: string]: any },) {
  return request<ApiListResp>('/service/api/list', {
    method: 'POST', headers: {
      ...headers,
    }, data: {
      page: params.current, ...params,
    }, ...(options || {}),
  });
}
/** 删除api  POST /service/api/delete*/
export async function deleteApi(options?: { [key: string]: any }) {
  return request<BaseResp>('/service/api/delete', {
    method: 'POST', headers: {
      ...headers,
    },
    data: {
      ...options,
    },
    ...(options || {}),
  });
}


