import {BaseResp} from "@/services/typings";
export type Sms = {
  id?: number;
  noticeCount?: number;
  usedNotice?: number;
}
export type SmsSend = {
  id?: number;
  createdAt?: string;
  status?: number;
  mobile?: string;
  code?: string;
  bizId?: string;
  userType?: string;
  content?: string;
  templates?: string;
  noticeCount?: string;
}


export type SmsResp={
  data?:  Sms;
  baseResp?: BaseResp;
}
export type SmsSendResp={
  data?:  SmsSend;
  baseResp?: BaseResp;
}
export type  SmsSendListResp = {
  data?:  SmsSend[];
  baseResp?: BaseResp;
}






