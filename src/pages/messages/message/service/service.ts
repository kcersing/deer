import {request} from '@umijs/max';
import {ItemListResp, ItemResp, Item, MessagesSendListResp, MessagesListResp, SmsSendListResp} from "./data";
import {TreeResp,Tree,BaseResp} from "@/services/typings";

const headers = {
  'Content-Type': 'application/json', Authorization: 'Bearer ' + sessionStorage.getItem('token') || '',
};

/** 获取 sms  POST /service/message/sms */
export async function getSms(options?: { [key: string]: any }) {
  console.log(options)
  return request<SmsResp>('/service/message/sms', {
    method: 'POST', headers: {
      ...headers,
    },
    data: {
      ...options,
    },
    ...(options || {}),
  });
}
/** sms  send 列表 POST /service/message/sms-send-list */
export async function getSmsSendList(params: {
  // query
  /** 当前的页码 */
  current?: number; /** 页面的容量 */
  pageSize?: number; keywords?: string;

}, options?: { [key: string]: any },) {
  return request<SmsSendListResp>('/service/message/sms-send-list', {
    method: 'POST', headers: {
      ...headers,
    },
    data: {
      page: params.current, ...params,
    }, ...(options || {}),
  });
}
/** send member messages  POST /service/message/send-member-messages */
export async function SendMemberMessages(options?: { [key: string]: any }) {
  return request<BaseResp>('/service/message/send-member-messages', {
    method: 'POST', headers: {
      ...headers,
    },
    data: {
      ...options,
    },
    ...(options || {}),
  });
}
/** send user messages POST /service/message/send-user-messages */
export async function SendUserMessages(options?: { [key: string]: any }) {
  return request<BaseResp>('/service/message/send-user-messages', {
    method: 'POST', headers: {
      ...headers,
    },
    data: {
      ...options,
    },
    ...(options || {}),
  });
}
  /** Messages列表 POST /service/message/messages-list */
  export async function getMessagesList(params: {
    // query
    /** 当前的页码 */
    current?: number; /** 页面的容量 */
    pageSize?: number; keywords?: string;

  }, options?: { [key: string]: any },) {
    return request<MessagesListResp>('/service/message/messages-list', {
      method: 'POST', headers: {
        ...headers,
      },
      data: {
        page: params.current, ...params,
      }, ...(options || {}),
    });
  }

  /** 发送记录 POST /service/message/messages-send-list */
  export async function getMessagesSendList(params: {
    // query
    /** 当前的页码 */
    current?: number; /** 页面的容量 */
    pageSize?: number; keywords?: string;

  }, options?: { [key: string]: any },) {
    return request<MessagesSendListResp>('/service/message/messages-send-list', {
      method: 'POST', headers: {
        ...headers,
      },
      data: {
        page: params.current, ...params,
      }, ...(options || {}),
    });
  }
