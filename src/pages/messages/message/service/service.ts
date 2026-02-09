import {request} from '@umijs/max';
import {ItemListResp, ItemResp, Item, MessagesSendListResp, MessagesListResp, SmsSendListResp} from "./data";
import {TreeResp,Tree,BaseResp} from "@/services/typings";

const headers = {
  'Content-Type': 'application/json', Authorization: 'Bearer ' + sessionStorage.getItem('token') || '',
};


/** send messages POST /service/message/send-messages */
export async function SendMessages(options?: { [key: string]: any }) {
  return request<BaseResp>('/service/message/send-messages', {
    method: 'POST', headers: {
      ...headers,
    },
    data: {
      ...options,
    },
    ...(options || {}),
  });
}
  /** Messages列表 POST /service/message/list */
  export async function getMessagesList(params: {
    // query
    /** 当前的页码 */
    current?: number; /** 页面的容量 */
    pageSize?: number; keywords?: string;

  }, options?: { [key: string]: any },) {
    return request<MessagesListResp>('/service/message/list', {
      method: 'POST', headers: {
        ...headers,
      },
      data: {
        page: params.current, ...params,
      }, ...(options || {}),
    });
  }

  /** 发送记录 POST /service/message/send/list */
  export async function getMessagesSendList(params: {
    // query
    /** 当前的页码 */
    current?: number; /** 页面的容量 */
    pageSize?: number; keywords?: string;

  }, options?: { [key: string]: any },) {
    return request<MessagesSendListResp>('/service/message/send/list', {
      method: 'POST', headers: {
        ...headers,
      },
      data: {
        page: params.current, ...params,
      }, ...(options || {}),
    });
  }

/** 删除item  POST /service/message/delete */
export async function deleteMessages(options?: { [key: string]: any }) {
  return request<BaseResp>('/service/message/delete', {
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

/** 列表 POST /service/message/types */
export async function getMessagesTypes() {
  return request<MessagesSendListResp>('/service/message/types', {
    method: 'POST', headers: {
      ...headers,
    },
  });
}


/** 更新Api  POST /service/message/send/update */
export async function updateSend(options?: { [key: string]: any }) {

  return request<ApiResp>('/service/message/send/update', {
    method: 'POST', headers: {
      ...headers,
    },
    data: {
      ...options,
    },
    ...(options || {}),
  });
}
