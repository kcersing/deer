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

export type MessagesSend = {
  id?: number;
  createdAt?: string;
  status?: string;
  receivedAt?: string;
  readAt?: string;
  type?: string;
  content?: string;
  messagesId?: number;
  fromUserId?: number;
  fromUserName?: string;
}

export type Messages = {
  id?: number;
  createdAt?: string;
  status?: string;
  createdId?: number;
  createdName?: number;
  type?: string;
  title?: string;
  content?: string;

  categoryId?: number;
  categoryName?: string;
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



export type MessagesListResp={
  data?: Messages[];
  baseResp?: BaseResp;
}
export type MessagesSendListResp = {
  data?: MessagesSend[];
  baseResp?: BaseResp;
}




