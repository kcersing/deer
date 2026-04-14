import {BaseResp} from "@/services/typings";



export type MessagesSend = {
  id?: number;
  status?: string;
  receivedAt?: string;
  readAt?: string;
  type?: string;
  content?: string;
  title?: string;
  fromUserId?: number;
  fromUserName?: string;
}

export type MessagesSendListResp = {
  data?: MessagesSend[];
  baseResp?: BaseResp;
}




