//日志信息
import {Dict} from "@/pages/system/dict/service/data";

export type Log = {
  type?: number;
  method?: string;
  api?: string;
  success?: bool;
  reqContent?: string;
  respContent?: string;
  ip?: string;
  userAgent?: string;
  operatorsr?: string;
  time?: number;
  createdAt?: number;
  updatedAt?: string;
  identity?: number;
  id?: number;
}

export type DeleteLogReq = {
  endAt?: string;
  startAt?: string;
}
export type LogListResp ={
  data?: Log[];
  /** 列表的内容总数 */
  baseResp?: BaseResp;
}
