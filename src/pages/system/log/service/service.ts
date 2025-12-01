// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';
import { headers } from '@/services/ant-design-pro/utils';

import { LogListResp,DeleteLogReq} from "./data";
import {BaseResp} from "@/services/typings";




/** 获取日志列表 POST /service/log/list */
export async function getLogList(options?: { [key: string]: any }) {
  return request<LogListResp>('/service/log/list', {
    method: 'POST',
    headers: {
      ...headers,
    },
    ...(options || {}),
  });
}
/** 删除 PSOT /service/log/delete */
export async function deleteLog(req: DeleteLogReq,options?: { [key: string]: any }) {
  return request<BaseResp>('/service/log/delete', {
    method: 'PSOT',
    headers: {
      ...headers,
    },
    ...(options || {}),
  });
}
