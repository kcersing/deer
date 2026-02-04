// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';
import { headers } from '@/services/ant-design-pro/utils';

import { LogListResp,DeleteLogReq} from "./data";
import {BaseResp} from "@/services/typings";




/** 获取日志列表 POST /service/logs/list */
export async function getLogList(options?: { [key: string]: any }) {
  return request<LogListResp>('/service/logs/list', {
    method: 'POST',
    headers: {
      ...headers,
    },
    data: {
      ...options,
    },
    ...(options || {}),
  });
}
/** 删除 POST /service/logs/delete */
export async function deleteLog(req: DeleteLogReq,options?: { [key: string]: any }) {
  return request<BaseResp>('/service/logs/delete', {
    method: 'POST',
    headers: {
      ...headers,
    },
    data: {
      ...options,
    },
    ...(options || {}),
  });
}
