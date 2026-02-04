import {BaseResp} from "@/services/typings";

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

export type MessagesListResp={
  data?: Messages[];
  baseResp?: BaseResp;
}
export type MessagesSendListResp = {
  data?: MessagesSend[];
  baseResp?: BaseResp;
}




